package backend

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
