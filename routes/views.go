package routes

import (
	"net/http"

	"tofus-blog/backend"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = backend.GetDB()
}

func homeView(w http.ResponseWriter, r *http.Request) {
	var posts = []backend.Post{}
	query := db.Order("updated_at").Find(&posts)

	if query.Error != nil {
		JSONResponse(w, 103, query.Error, "Error fetching records")
	}

	renderTemplate(w, r, "post_list.html", map[string]any{"posts": posts, "data": 123})
}
