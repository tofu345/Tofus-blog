package src

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ! Have post likes be positioned with grid not padding
// ? Change db error messages
// ? Expand post on post list page instead of redirect
// ! use varaibles in css
// ! implement something similar to messages in django

func homeView(w http.ResponseWriter, r *http.Request) {
	objects := []Post{}
	err := db.Find(&objects).Error
	if err != nil {
		ErrorResponse(w, r, err, nil)
		return
	}

	// ? Reverse list on front end instead
	j := len(objects) - 1
	for i := 0; i < j; i++ {
		objects[i], objects[j] = objects[j], objects[i]
		j--
	}

	RenderTemplate(w, r, "posts/post_list.html",
		map[string]any{"posts": objects}, NewTemplateConfig())
}

func NotFound404Handler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "error.html",
		map[string]any{
			"data": fmt.Sprintf("The page %v%v was not found", r.Host, r.URL),
			"err":  "404 Not Found",
		}, &TemplateConfig{})
}

func postDetailView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	post, err := GetPostBySlug(slug)
	if err != nil {
		ErrorResponse(w, r, errors.New("Post Not Found"),
			fmt.Sprintf("No post with slug %v was found", slug))
		return
	}

	RenderTemplate(w, r, "posts/post_detail.html",
		map[string]any{"post": post}, NewTemplateConfig())
}

func loginView(w http.ResponseWriter, r *http.Request) {
	_, err := getUserFromRequest(w, r)
	if err == nil {
		// User logged in - redirect to home
		http.Redirect(w, r, baseUrl, http.StatusSeeOther)
		return
	}

	RenderTemplate(w, r, "login.html", map[string]any{}, &TemplateConfig{})
}

func signUpView(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, 103, nil, "Not implemented")
}
