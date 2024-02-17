package routes

import (
	"ecommerce_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// Product Item Routes
	router := gin.Default()

	// Routes
	router.GET("/items", controllers.GetItems)

	// Item ID Routes
	router.GET("/items/:id", controllers.ItemHandler)
	router.POST("/items", controllers.ItemHandler)
	router.PUT("/items/:id", controllers.ItemHandler)
	router.DELETE("/items/:id", controllers.ItemHandler)

	// Base Route
	router.GET("/", controllers.BaseHandler)

	return router
}
