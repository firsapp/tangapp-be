package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID, email, name string, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(24 * time.Hour).Unix(), //	It's a syntax to set token expiration 24 hours from now
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
