package store

import (
	"database/sql"
	"errors"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/anindita02/books-api-server/pkg/common"
)

// BookStore represents the store for managing books
type BookStore struct {
	db *sql.DB
}

type Store interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]*models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

var bookStore *BookStore

// NewBookStore creates a new instance of BookStore
func NewBookStore() *BookStore {
	db := common.Connect()
	if bookStore == nil {
		bookStore = &BookStore{db: db}
	}
	return bookStore
}

// Close closes the database connection
func (s *BookStore) Close() error {
	return s.db.Close()
}

// CreateBook creates a new book record
func (s *BookStore) CreateBook(book *models.Book) error {
	stmt, err := s.db.Prepare("INSERT INTO books(title, author) VALUES($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(book.Name, book.Author, book.Year, book.Language).Scan(&book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetBooks retrieves all book records
func (s *BookStore) GetBooks() ([]*models.Book, error) {
	// Fetch all books from the database
	rows, err := s.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Year, &book.Language)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

// GetBookByID retrieves a book record by its ID
func (s *BookStore) GetBookByID(id int) (*models.Book, error) {
	stmt, err := s.db.Prepare("SELECT id, title, author FROM books WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var book models.Book
	err = stmt.QueryRow(id).Scan(&book.ID, &book.Name, &book.Author, &book.Year, &book.Language)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no record is found
		}
		return nil, err
	}

	return &book, nil
}

// UpdateBook updates a book record
func (s *BookStore) UpdateBook(book *models.Book) error {
	stmt, err := s.db.Prepare("UPDATE books SET title = $1, author = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Name, book.Author, book.Year, book.Language, book.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBook deletes a book record by its ID
func (s *BookStore) DeleteBook(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
