package handlers

import (
	"net/http"

	"github.com/anindita02/books-api-server/pkg/app/models"
	"github.com/anindita02/books-api-server/pkg/app/services"
	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// AuthenticateUser authenticates a user and generates an access token.
// @Summary Authenticate user and generate access token
// @Description Authenticates a user with the provided credentials and generates an access token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.User true "User credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var credentials *models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	token, err := h.userService.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
