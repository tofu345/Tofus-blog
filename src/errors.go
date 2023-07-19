package src

import "errors"

var (
	ObjectNotFound = errors.New("Object not found")
	InvalidToken   = errors.New("Invalid token")
	TokenExpired   = errors.New("Token expired")
	LoginError     = errors.New("Incorrect username or password")
	Unauthorized   = errors.New("Unauthorized")
)
