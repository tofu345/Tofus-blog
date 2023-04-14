package user

import (
	"tofs-blog/backend"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = backend.GetDB()
}

func Create(user *backend.User) error {
	result := db.Create(&user)
	return result.Error
}

func Errors(user *backend.User) map[string]string {
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
