package src

import (
	"fmt"
	"net/http"
)

func homeView(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	err := db.Find(&posts).Error
	if err != nil {
		RenderErrorPage(w, r, err, nil)
		return
	}

	// user, err := getUserFromRequest(w, r)
	// if err != nil {
	// 	http.Redirect(w, r, baseUrl, http.StatusSeeOther)
	// 	return
	// }

	RenderTemplate(w, r, "posts/post_list.html",
		map[string]any{"posts": posts}, &TemplateConfig{NavbarShown: true})
}

func NotFound404Handler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "error.html",
		map[string]any{
			"data": fmt.Sprintf("The page %v%v was not found", r.Host, r.URL),
			"err":  "404 Not Found",
		}, &TemplateConfig{})
}

func postDetailView(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		JSONError(w, err)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		RenderErrorPage(w, r, ErrObjectNotFound, fmt.Sprintf("No post with id %v was found", id))
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

	fmt.Println(err)

	RenderTemplate(w, r, "login.html", map[string]any{}, &TemplateConfig{})
}

func signUpView(w http.ResponseWriter, r *http.Request) {
	Response(w, 400, dict{"message": "Not implemented"})
}
