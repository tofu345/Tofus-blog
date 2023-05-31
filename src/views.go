package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeView(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	err := db.Find(&posts).Error
	if err != nil {
		RenderErrorPage(w, r, err, nil)
		return
	}

	json, err := json.Marshal(posts)
	if err != nil {
		RenderErrorPage(w, r, err, nil)
		return
	}

	RenderTemplate(w, r, "posts/post_list.html",
		map[string]any{"posts": string(json), "postLen": len(posts)}, &TemplateConfig{NavbarShown: true})
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
		RenderErrorPage(w, r, errors.New("Post Not Found"),
			fmt.Sprintf("No post with slug %v was found", slug))
		return
	}

	RenderTemplate(w, r, "posts/post_detail.html",
		map[string]any{"post": post}, &TemplateConfig{NavbarShown: true})
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
