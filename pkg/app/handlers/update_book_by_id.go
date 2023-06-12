package handlers

import (
	"net/http"
	"strconv"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

// UpdateBookByID updates a book.
// @Summary Updates a book
// @Description Updates a book with the provided details.
// @Tags Books
// @ID                     books-update
// @Security               ApiKeyAuth
// @Param  request         body     models.Book  true   "update books request"
// @Param  id              path     string                                  true   "postgresql id of the book"
// @Accept                 json
// @Produce                json
// @Param  authorization   header   string      true   "authorization token"
// @Success 200 string    "Book updated successfully"
// @Failure 400 {object}   map[string]interface{}
// @Failure 500 {object}   map[string]interface{}
// @Router /api/update_book/{id} [put]
func (bH *BookHandler) UpdateBookByID(c *gin.Context) {
	idStr := c.Param("id") // Get the book ID from the request parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	var book *models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	book.ID = id
	err = bH.service.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book entry"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
