package routes

import (
	"errors"
	"fmt"
	"net/http"

	"tofs-blog/posts"

	"github.com/gorilla/mux"
)

func homeView(w http.ResponseWriter, r *http.Request) {
	objects, err := posts.GetAll()
	if err != nil {
		ErrorResponse(w, r, err, nil)
	}

	// ? Reverse list on front end instead
	j := len(objects) - 1
	for i := 0; i < j; i++ {
		objects[i], objects[j] = objects[j], objects[i]
		j--
	}

	RenderTemplate(w, r, "posts/post_list.html",
		map[string]any{"posts": objects, "data": 123}, NewTemplateConfig())
}

func NotFound404Handler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "error.html",
		map[string]any{"data": fmt.Sprintf("The page %v was not found", r.URL), "err": "404 Not Found"}, &TemplateConfig{})
}

func postDetailView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	post, e := posts.GetBySlug(slug)
	if e != nil {
		ErrorResponse(w, r, errors.New("Post Not Found"), fmt.Sprintf("No post with slug %v was found", slug))
		return
	}

	RenderTemplate(w, r, "posts/post_detail.html",
		map[string]any{"post": post}, NewTemplateConfig())
}
