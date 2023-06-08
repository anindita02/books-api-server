package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a new JWT token
func GenerateToken(userID int64, secretKey string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expirationTime).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
