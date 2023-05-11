package src

import (
	"errors"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID       int       `gorm:"primarykey" json:"id"`
	Title    string    `json:"title"`
	Slug     string    `json:"slug"`
	Body     string    `json:"body"`
	Author   string    `json:"author"`
	Views    int       `json:"views"`
	Likes    uint64    `json:"likes"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ID"`
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

type User struct {
	ID              int       `gorm:"primarykey" json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username" gorm:"unique"`
	Password        string    `json:"password"`
	Email           string    `json:"email" gorm:"unique"`
	AccessToken     string    `json:"-" gorm:"unique"` // Exclude from JSON serialization
	TokenExpiryDate time.Time `json:"token_expiry_date"`
	BaseModel
}

func getUserByToken(token string) (User, error) {
	var user User
	err := db.First(&user, "access_token = ?", token).Error
	if err != nil {
		if err.Error() == RecordNotFound {
			return User{}, errors.New(NoTokenFound)
		}
		return User{}, err
	}

	return user, err
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
