package handlers

import (
	"net/http"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

func (bH *BookHandler) CreateBooks(c *gin.Context) {
	var books *models.Book
	if err := c.BindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := bH.service.CreateBook(books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create books entry"})
		return
	}

	c.JSON(http.StatusOK, &books)
}
