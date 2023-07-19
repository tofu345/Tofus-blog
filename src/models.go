package src

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID         int       `gorm:"primarykey" json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Body       string    `json:"body"`
	AuthorID   int       `json:"author_id"`
	AuthorName string    `json:"author"`
	Views      int       `json:"views"`
	Likes      uint64    `json:"likes"`
	Comments   []Comment `json:"comments" gorm:"foreignKey:ID"`
	BaseModel
}

func GetPostBySlug(slug string) (Post, error) {
	var post Post
	result := db.First(&post, "slug = ?", slug)
	return post, result.Error
}

func (post *Post) Errors() map[string]string {
	errors := map[string]string{}

	if post.Title == "" {
		errors["title"] = "This field is required"
	}

	if post.Body == "" {
		errors["body"] = "This field is required"
	}

	return errors
}

type Comment struct {
	BaseModel
	ID      int    `gorm:"primarykey" json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

type Permission struct {
	ID          int    `gorm:"primarykey"`
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}

type User struct {
	ID              int          `gorm:"primarykey" json:"id"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	Username        string       `json:"username" gorm:"unique"`
	Password        string       `json:"-"`
	Email           string       `json:"email" gorm:"unique"`
	TokenExpiryDate time.Time    `json:"token_expiry_date"`
	UserPerms       []Permission `json:"permissions" gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsAdmin         bool         `json:"is_admin"`
	BaseModel
}

func (user *User) Errors() map[string]string {
	errors := map[string]string{}

	if user.Username == "" {
		errors["username"] = "This field is required"
	}

	if user.FirstName == "" {
		errors["first_name"] = "This field is required"
	}

	if user.Password == "" {
		errors["password"] = "This field is required"
	}

	if user.Email == "" {
		errors["email"] = "This field is required"
	}

	return errors
}
