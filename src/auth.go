package src

import (
	"net/http"
	"strings"
)

func userFromBearer(w http.ResponseWriter, r *http.Request) (User, error) {
	token := r.Header.Get("Authorization")
	token = strings.Split(token, " ")[1]

	return userFromToken(token)
}

func userFromToken(token string) (User, error) {
	payload, err := decodeToken(token)
	if err != nil {
		return User{}, ErrInvalidToken
	}

	username := payload["username"]

	switch username := username.(type) {
	case string:
		var user User
		err := db.First(&user, "username = ?", username).Error
		if err != nil {
			return User{}, err
		}

		return user, nil
	default:
		return User{}, ErrInvalidToken
	}
}
