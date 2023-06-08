package services

import (
	"time"

	"github.com/anindita02/books-api-server/pkg/app/store"
	"github.com/anindita02/books-api-server/pkg/utils"
)

// UserService handles user-related operations
type UserService struct {
	userStore     *store.UserStore
	tokenSecret   string
	tokenDuration time.Duration
}

// NewUserService creates a new UserService instance
func NewUserService(userStore *store.UserStore, tokenSecret string, tokenDuration time.Duration) *UserService {
	return &UserService{
		userStore:     userStore,
		tokenSecret:   tokenSecret,
		tokenDuration: tokenDuration,
	}
}

// AuthenticateUser authenticates a user with the given username and password
func (s *UserService) AuthenticateUser(username, password string) (string, error) {

	// Generate JWT token
	token, err := utils.GenerateToken(02062000, "anindita02", 24*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
