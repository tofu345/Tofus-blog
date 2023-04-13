package routes

import (
	"net/http"

	"tofs-blog/backend"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = backend.GetDB()
}

func homeView(w http.ResponseWriter, r *http.Request) {
	var posts = []backend.Post{}
	query := db.Find(&posts).Order("updated_at")

	// Reverse the list
	// Couldnt think of any better way to sort by last updated first
	// but to do it on the frontend
	j := len(posts) - 1
	for i := 0; i < j; i++ {
		posts[i], posts[j] = posts[j], posts[i]
		j--
	}

	if query.Error != nil {
		JSONResponse(w, 103, query.Error, "Error fetching records")
	}

	RenderTemplate(w, r, "posts/post_list.html",
		map[string]any{"posts": posts, "data": 123}, DefaultTemplateConfig)
}

func postDetailView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	post, e := backend.GetPostBySlug(slug)
	if e != nil {
		JSONResponse(w, 103, e, "Post Not Found")
		return
	}

	RenderTemplate(w, r, "posts/post_detail.html", map[string]any{"post": post}, DefaultTemplateConfig)
}
