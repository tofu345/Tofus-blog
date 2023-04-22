package db

import (
	"time"
)

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

func GetUser(user *User) error {
	result := DB.First(&user, "email = ?", user.Email)
	return result.Error
}

func GetUserByToken(token string) (User, error) {
	var user User
	result := DB.First(&user, "access_token = ?", token)
	return user, result.Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	result := DB.First(&user, "email = ?", email)
	return &user, result.Error
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
