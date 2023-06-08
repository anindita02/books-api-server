package handlers

import (
	"net/http"
	"strconv"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

// updateBook updates a book by its ID
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
