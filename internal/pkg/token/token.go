package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("password"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
