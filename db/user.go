package db

type User struct {
	BaseModel
	ID          int    `gorm:"primarykey" json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"unique"`
	AccessToken string `json:"-" gorm:"unique"`
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
