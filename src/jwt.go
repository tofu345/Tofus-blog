package src

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func newAccessToken(username string) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	time_now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "my-auth-server",
		"iat":      time_now.Unix(),
		"exp":      time_now.Add(time.Hour).Unix(),
		"username": username,
	})
	return t.SignedString(key)
}

func newRefreshToken(username string) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	time_now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "my-auth-server",
		"iat":      time_now.Unix(),
		"exp":      time_now.Add(time.Hour).Unix(),
		"ref":      true,
		"username": username,
	})
	return t.SignedString(key)
}

func decodeToken(tokenData string) (map[string]any, error) {
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return map[string]any{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return map[string]any{}, err
	}
}
