package handlers

import (
	"net/http"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/gin-gonic/gin"
)

// CreateBook creates a new book.
// @Summary Create a new book
// @Description Creates a new book with the provided details.
// @Tags Books
// @ID                     books-create
// @Security               ApiKeyAuth
// @Accept                 json
// @Produce                json
// @Param  request         body     models.Book true  "create book request"
// @Param  authorization   header   string      true   "authorization token"
// @Success 200 {object}    models.Book
// @Failure 400 {object}   map[string]interface{}
// @Failure 500 {object}   map[string]interface{}
// @Router /api/book [post]
func (bH *BookHandler) CreateBook(c *gin.Context) {
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
