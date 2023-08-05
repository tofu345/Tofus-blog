package src

import "errors"

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrInvalidToken   = errors.New("invalid token")
	ErrTokenExpired   = errors.New("token expired")
	ErrLoginError     = errors.New("incorrect username or password")
	ErrUnauthorized   = errors.New("Unauthorized")
)
