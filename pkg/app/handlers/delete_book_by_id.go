package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (bH *BookHandler) DeleteBookByID(c *gin.Context) {
	idStr := c.Param("id") // Get the book ID from the request parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = bH.service.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
