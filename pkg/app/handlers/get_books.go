package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getBooks retrieves all books
func (bH *BookHandler) GetBooks(c *gin.Context) {
	book, err := bH.service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}

	c.JSON(http.StatusOK, book)
}
