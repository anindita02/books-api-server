package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// VerifyToken verifies and parses a JWT token
func VerifyToken(tokenString, secretKey string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Claims represents the custom claims for JWT token
type Claims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

// AuthMiddleware is a middleware to authenticate requests
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Extract the token from the "Bearer <token>" format
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		tokenString = tokenParts[1]

		// Validate the token
		claims, err := VerifyToken(tokenString, "your-secret-key")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set the authenticated user ID in the context for further processing
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
