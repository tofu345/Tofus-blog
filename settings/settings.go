package settings

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT = map[string]any{
		"ISSUER":                 "sherlock",
		"ACCESS_TOKEN_LIFETIME":  time.Hour,
		"REFRESH_TOKEN_LIFETIME": time.Hour * 24,
		"ALGORITHM":              jwt.SigningMethodHS256,
	}
)
