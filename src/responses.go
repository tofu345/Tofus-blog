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

type dict map[string]any

type TemplateConfig struct {
	NavbarShown bool
}

const (
	baseUrl = "http://localhost:8005"
)

func idParam(req *http.Request) (int, error) {
	vars := mux.Vars(req)
	return strconv.Atoi(vars["id"])
}

func getUserFromRequest(w http.ResponseWriter, r *http.Request) (User, error) {
	token, err := r.Cookie("access")
	if err != nil {
		return User{}, ErrInvalidToken
	}

	return userFromToken(token.Value)
}

func getUserFromRequestAndRedirect(w http.ResponseWriter, r *http.Request) (User, error) {
	user, err := getUserFromRequest(w, r)
	if err == nil {
		return user, err
	}

	if errors.Is(err, ErrInvalidToken) || errors.Is(err, ErrTokenExpired) {
		url := fmt.Sprintf("%v/login?next=%v", baseUrl, r.URL)
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		RenderErrorPage(w, r, err, "Error")
	}

	// fmt.Println(err.Error())

	return user, err
}

func Response(w http.ResponseWriter, code int, data dict) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	if code == 200 {
		data["responseCode"] = 100
	} else {
		data["responseCode"] = 103
	}

	json.NewEncoder(w).Encode(data)
}

func ParseError(err error) string {
	str := err.Error()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		str = "Object not found"
	case strings.HasPrefix(str, "UNIQUE constraint failed: "):
		str = strings.Split(str, ": ")[1] + " is already in use"
	}

	return str
}

func JSONError(w http.ResponseWriter, err any) {
	switch err := err.(type) {
	case error:
		Response(w, 400, dict{"message": "An Error Occured", "error": ParseError(err)})
		return
	}
	Response(w, 400, dict{"message": "An Error Occured", "error": err})
}

func JSONDecode(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func RenderErrorPage(w http.ResponseWriter, r *http.Request, err error, data any) {
	log.Println("Error: " + err.Error())

	RenderTemplate(w, r, "error.html",
		map[string]any{"data": data, "err": err.Error()}, &TemplateConfig{})
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
