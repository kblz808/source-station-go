package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userID interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
