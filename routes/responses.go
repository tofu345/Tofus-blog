package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Response struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
	Data         any    `json:"data"`
}

type TemplateConfig struct {
	NavbarShown     bool
	BackgroundShown bool
}

func NewTemplateConfig() *TemplateConfig {
	return &TemplateConfig{
		NavbarShown:     true,
		BackgroundShown: true,
	}
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

// Renders error page
func ErrorResponse(w http.ResponseWriter, r *http.Request, err error, data any) {
	log.Println(err)

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
		"fmt_likes": func(likes int) string {
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
