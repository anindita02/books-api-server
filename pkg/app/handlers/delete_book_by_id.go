package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteBookByID deletes a book by its ID
// @Summary Deletes a book by its ID
// @Description delete book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param  id              path     string                                  true   "postgresql id of the book"
// @Success 200 string   "Book deleted successfully"
// @Security ApiKeyAuth
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/delete_book/{id} [delete]
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
