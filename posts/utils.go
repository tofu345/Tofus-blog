package posts

import (
	"tofs-blog/backend"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = backend.GetDB()
}

func GetAll() ([]backend.Post, error) {
	var posts = []backend.Post{}
	query := db.Find(&posts)
	return posts, query.Error
}

func GetById(id int) (backend.Post, error) {
	var post = backend.Post{ID: id}
	result := db.First(&post)
	return post, result.Error
}

func GetBySlug(slug string) (backend.Post, error) {
	var post backend.Post
	result := db.First(&post, "slug = ?", slug)
	return post, result.Error
}

func Delete(id int) {
	db.Exec("DELETE FROM posts WHERE id = ?", id)
}

func Create(post *backend.Post) error {
	result := db.Create(&post)
	return result.Error
}

func Update(post *backend.Post) error {
	result := db.Save(&post)
	return result.Error
}

func Errors(post backend.Post) map[string]string {
	errors := map[string]string{}

	if post.Title == "" {
		errors["title"] = "This field is required"
	}

	if post.Body == "" {
		errors["body"] = "This field is required"
	}

	if post.Author == "" {
		errors["author"] = "This field is required"
	}

	return errors
}
