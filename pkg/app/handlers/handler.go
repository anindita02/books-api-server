package handlers

import "github.com/anindita02/books-api-server/pkg/app/services"

type BookHandler struct {
	service *services.BookService
}

var bookHandler *BookHandler

// NewBookHandler creates a new instance of BookHandler
func NewBookHandler(bookService *services.BookService) *BookHandler {
	if bookHandler == nil {
		bookHandler = &BookHandler{service: bookService}
	}
	return bookHandler
}
