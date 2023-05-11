package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
	Data         any    `json:"data"`
}

type TemplateConfig struct {
	NavbarShown bool
}

const (
	baseUrl = "http://localhost:8005"

	InvalidURL      = "Invalid URL"
	InvalidPOSTData = "Invalid POST Data"
	InvalidData     = "Invalid Data"

	NoTokenFound = "Invalid Token"
	TokenExpired = "Token Expired"

	RecordNotFound = "record not found"
	LoginError     = "Incorrect username or password"
)

func getIdFromRequest(req *http.Request) (int, error) {
	vars := mux.Vars(req)
	return strconv.Atoi(vars["id"])
}

// todo: unnecessary right now
// func formatDbError(errPtr *error) {
// 	err := *errPtr
// 	switch err.Error() {
// 	case "record not found":

// 	}
// }

func getUserFromRequestApi(w http.ResponseWriter, r *http.Request) (User, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return User{}, errors.New(NoTokenFound)
	}

	token = strings.Split(token, " ")[1]
	user, err := getUserByToken(token)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Uses user token to get logged in user
// returns empty user object if not logged in
func getUserFromRequest(w http.ResponseWriter, r *http.Request) (User, error) {
	token, err := r.Cookie("token")
	if err != nil {
		return User{}, errors.New(NoTokenFound)
	}

	user, err := getUserByToken(token.Value)
	if err != nil {
		return User{}, err
	}

	currentTime := time.Now()
	if user.TokenExpiryDate.Before(currentTime) {
		return User{}, errors.New(TokenExpired)
	}

	return user, nil
}

// redirects to login page if not logged in and renders error page if user not found or any other error
func getUserFromRequestAndRedirect(w http.ResponseWriter, r *http.Request) (User, error) {
	user, err := getUserFromRequest(w, r)
	if err == nil {
		return user, err
	}

	if err.Error() == NoTokenFound || err.Error() == TokenExpired {
		url := fmt.Sprintf("%v/login?next=%v", baseUrl, r.URL)
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		ErrorResponse(w, r, err, "Error")
	}

	// fmt.Println(err.Error())

	return user, err
}

// responseCode: 103 - Bad, 100 - Good
func JSONResponse(w http.ResponseWriter, responseCode int, data any, message string) {
	w.Header().Set("Content-type", "application/json")

	if responseCode == 103 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(Response{
		ResponseCode: responseCode,
		Data:         data,
		Message:      message,
	})
}

func JSONError(w http.ResponseWriter, err string) {
	if err == RecordNotFound {
		err = "Object Not Found"
	} else if strings.HasPrefix(err, "UNIQUE constraint failed: ") {
		attr := strings.Split(strings.Split(err, ": ")[1], ".")[1]
		err = strings.Title(attr + " is already in use")
	}

	JSONResponse(w, 103, err, "Error")
}

// Renders error page
func ErrorResponse(w http.ResponseWriter, r *http.Request, err error, data any) {
	log.Println("Error: " + err.Error())

	RenderTemplate(w, r, "error.html",
		map[string]any{"data": data, "err": err.Error()}, &TemplateConfig{})
}

// Decodes request body into data
//
// data should be a pointer
func JSONDecode(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func parseFiles(funcs template.FuncMap, filenames ...string) (*template.Template, error) {
	return template.New(filepath.Base(filenames[0])).Funcs(funcs).ParseFiles(filenames...)
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, pathToFile string, data map[string]any, config *TemplateConfig) {
	data["template_config"] = *config
	// Add user object if not exists
	if _, exists := data["user"]; !exists {
		data["user"], _ = getUserFromRequest(w, r)
	}

	baseTemplateDir := "templates"

	lp := filepath.Join(baseTemplateDir, "layout.html")
	fp := filepath.Join(baseTemplateDir, filepath.Clean(pathToFile))

	// fmt.Println(pathToFile)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	// Template Functions
	funcs := template.FuncMap{
		"truncateStr": func(text string) string {
			max := 500
			if max > len(text) {
				return text
			}
			return text[:strings.LastIndex(text[:max], " ")] + "..."
		},
		"ls_or_eq": func(num1, num2 int) bool {
			return num1 <= num2
		},
		"gt": func(num1, num2 int) bool {
			return num1 > num2
		},
		"fmt_likes": func(likes uint64) string {
			// < 1k
			if likes <= 999 {
				return fmt.Sprint(likes)
			}

			// < 1m
			if likes <= 999999 {
				return fmt.Sprintf("%dK", likes/1000)
			}

			return fmt.Sprintf("%dM", likes/1000000)
		},
		"userLoggedIn": func(user User) bool {
			return user != User{}
		},
	}

	tmpl, err := parseFiles(funcs, lp, fp)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
