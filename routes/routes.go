package routes

import (
	"ecommerce_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Product Item Routes
	productRoutes := router.Group("/products")
	{
		// Get all procuts
		productRoutes.GET("/products", controllers.GetItems)

		// Item ID Routes
		productRoutes.GET("/products/:id", controllers.ItemHandler)
		productRoutes.POST("/products", controllers.ItemHandler)
		productRoutes.PUT("/products/:id", controllers.ItemHandler)
		productRoutes.DELETE("/products/:id", controllers.ItemHandler)
	}

	// Base Route
	router.GET("/", controllers.BaseHandler)

	return router
}
