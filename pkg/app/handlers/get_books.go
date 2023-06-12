package handlers

import (
	"net/http"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

// GetBooks gets a list of books
// @Summary Get a list of books
// @Description Retrieves a list of books.
// @Tags Books
// @ID                     books-list
// @Security               ApiKeyAuth
// @Accept                 json
// @Produce                json
// @Param  authorization   header   string      true   "authorization token"
// @Success 200 {array}    models.Book
// @Failure 500 {object}   map[string]interface{}
// @Router /api/list_books [get]
func (bH *BookHandler) GetBooks(c *gin.Context) {
	var book []*models.Book
	book, err := bH.service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}

	c.JSON(http.StatusOK, book)
}
