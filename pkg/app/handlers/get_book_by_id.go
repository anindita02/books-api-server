package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getBookByID retrieves a book by its ID
func (bH *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id") // Get the book ID from the request parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := bH.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book by ID"})
		return
	}

	c.JSON(http.StatusOK, book)
}
