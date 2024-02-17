package routes

import (
	"ecommerce_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Product Routes
	productRoutes := router.Group("/products")
	{
		// Base route (get all products)
		productRoutes.GET("/products", controllers.GetItems)

		// Path param routes (Create, Read, Update, Delete products)
		productRoutes.GET("/products/:id", controllers.ItemHandler)
		productRoutes.POST("/products", controllers.ItemHandler)
		productRoutes.PUT("/products/:id", controllers.ItemHandler)
		productRoutes.DELETE("/products/:id", controllers.ItemHandler)
	}

	// Root Route
	router.GET("/", controllers.BaseHandler)

	return router
}
