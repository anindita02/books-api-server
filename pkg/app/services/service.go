package services

import (
	"errors"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/anindita02/books-api-server/pkg/app/store"
)

// BookService represents the service for managing books
type BookService struct {
	store *store.BookStore
}

type Service interface {
	CreateBook(book *models.Book) error
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	GetBooks() ([]*models.Book, error)
	DeleteBook(id int) error
}

var bookService *BookService

// NewBookService creates a new instance of BookService
func NewBookService(bookStore *store.BookStore) *BookService {
	if bookService == nil {
		bookService = &BookService{store: bookStore}
	}
	return bookService
}

// CreateBook creates a new book record
func (s *BookService) CreateBook(book *models.Book) error {
	if book == nil {
		return errors.New("book is required")
	}

	return s.store.CreateBook(book)
}

// GetBookByID retrieves a book record by its ID
func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.store.GetBookByID(id)
}

// UpdateBook updates a book record
func (s *BookService) UpdateBook(book *models.Book) error {
	if book == nil {
		return errors.New("book is required")
	}
	return s.store.UpdateBook(book)
}

// GetBooks retrieves all books
func (s *BookService) GetBooks() ([]*models.Book, error) {
	return s.store.GetBooks()
}

// DeleteBook deletes a book record by its ID
func (s *BookService) DeleteBook(id int) error {
	return s.store.DeleteBook(id)
}
