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

	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
)

var (
	TokenInvalid = errors.New("Invalid Token")
	TokenExpired = errors.New("Token Expired")
	LoginError   = errors.New("Incorrect username or password")
	Unauthorized = errors.New("You do not have permission to perform this action")
)

func getIdFromRequest(req *http.Request) (int, error) {
	vars := mux.Vars(req)
	return strconv.Atoi(vars["id"])
}

func getUserFromRequestApi(w http.ResponseWriter, r *http.Request) (User, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return User{}, TokenInvalid
	}

	token = strings.Split(token, " ")[1]
	user, err := getUserByToken(token)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserFromRequest(w http.ResponseWriter, r *http.Request) (User, error) {
	token, err := r.Cookie("token")
	if err != nil {
		return User{}, TokenInvalid
	}

	user, err := getUserByToken(token.Value)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserFromRequestAndRedirect(w http.ResponseWriter, r *http.Request) (User, error) {
	user, err := getUserFromRequest(w, r)
	if err == nil {
		return user, err
	}

	if errors.Is(err, TokenInvalid) || errors.Is(err, TokenExpired) {
		url := fmt.Sprintf("%v/login?next=%v", baseUrl, r.URL)
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		RenderErrorPage(w, r, err, "Error")
	}

	// fmt.Println(err.Error())

	return user, err
}

// responseCode: 103 - Bad, 100 - Good
func JSONResponse(w http.ResponseWriter, responseCode int, data any, message string) {
	w.Header().Set("Content-type", "application/json")

	if responseCode == 103 {
		w.WriteHeader(http.StatusBadRequest)
	} else if responseCode == 100 {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(Response{
		ResponseCode: responseCode,
		Data:         data,
		Message:      message,
	})
}

func JSONError(w http.ResponseWriter, err error) {
	str := err.Error()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		str = "Object Not Found"
	} else if strings.HasPrefix(str, "UNIQUE constraint failed: ") {
		// UNIQUE constraint failed: posts.title -> Title is already in use
		attr := strings.Split(strings.Split(str, ": ")[1], ".")[1]
		str = strings.Title(attr + " is already in use")
	}

	JSONResponse(w, 103, str, "Error")
}

func RenderErrorPage(w http.ResponseWriter, r *http.Request, err error, data any) {
	log.Println("Error: " + err.Error())

	RenderTemplate(w, r, "error.html",
		map[string]any{"data": data, "err": err.Error()}, &TemplateConfig{})
}

// Decodes request body into data
//
// data must be a pointer
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
		"ls_or_eq": func(num1, num2 int) bool {
			return num1 <= num2
		},
		"gt": func(num1, num2 int) bool {
			return num1 > num2
		},
		"userLoggedIn": func(user User) bool {
			return user.ID != 0
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
