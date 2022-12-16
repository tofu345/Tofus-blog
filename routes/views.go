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
	renderTemplate(w, r, "welcome.html", nil)
}
