package app

import (
	"github.com/anindita02/books-api-server/pkg/app/handlers"
	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine, aH *handlers.BookHandler) {
	router.GET("/", aH.GetBooks)
	router.POST("/", aH.CreateBooks)
	router.DELETE("/:id", aH.DeleteBookByID)
	router.PUT("/:id", aH.UpdateBookByID)
	router.GET("/:id", aH.GetBookByID)
}
