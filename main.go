package config

import (
	"time"

	"github.com/anindita02/books-api-server/pkg/app/handlers"
	"github.com/anindita02/books-api-server/pkg/app/services"
	store "github.com/anindita02/books-api-server/pkg/app/store"
	"github.com/anindita02/books-api-server/pkg/common"
	"github.com/anindita02/books-api-server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	db := common.Connect()
	router := gin.Default()

	// Create a user service and handler
	userStore := store.NewUserStore(db)
	userService := services.NewUserService(userStore, "anindita02", time.Hour*24)
	userHandler := handlers.NewUserHandler(userService)

	// Public routes (no authentication required)
	router.POST("/login", userHandler.AuthenticateUser)

	// Create a new instance of the store layer
	bookStore := store.NewBookStore()
	// Create a new instance of the service layer
	bookService := services.NewBookService(bookStore)

	// Create a new instance of the handler layer
	bookHandler := handlers.NewBookHandler(bookService)

	// Protected routes (authentication required)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/list_books", bookHandler.GetBooks)
		protected.POST("/book", bookHandler.CreateBook)
		protected.PUT("/update_book/:id", bookHandler.UpdateBookByID)
		protected.DELETE("/delete_book/:id", bookHandler.DeleteBookByID)
		protected.GET("/:id", bookHandler.GetBookByID)
	}

	return router
}

func main() {

	// Setup the router and routes
	router := setupRouter()

	// Start the server
	router.Run(":8080")
}
