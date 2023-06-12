package handlers

import (
	"net/http"
	"strconv"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

// GetBookByID retrieves a book by its ID
// @Summary Retrieves a book by its ID
// @Description get book by ID
// @Tags books
// @ID                     book-by-id
// @Accept  json
// @Produce  json
// @Param  id              path     string                                  true   "postgresql id of the book"
// @Success 200 {object} models.Book
// @Security ApiKeyAuth
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/{id} [get]
func (bH *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id") // Get the book ID from the request parameter

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	var book *models.Book
	book, err = bH.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book by ID"})
		return
	}

	c.JSON(http.StatusOK, book)
}
