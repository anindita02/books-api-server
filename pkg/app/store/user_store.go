package store

import (
	"database/sql"
	"errors"

	"github.com/anindita02/books-api-server/pkg/app/models"
)

// UserStore handles user data operations
type UserStore struct {
	db *sql.DB
}

// NewUserStore creates a new UserStore instance
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// GetUserByUsername retrieves a user by their username from the database
func (s *UserStore) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	user := &models.User{}

	err := s.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
