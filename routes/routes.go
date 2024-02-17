package routes

import (
	"ecommerce_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apiRoutes := router.Group("/api")

	v1Routes := apiRoutes.Group("/v1")
	{

		// Root Route
		router.GET("/", controllers.BaseHandler)

		// Product Routes
		productRoutes := v1Routes.Group("/products")
		{
			// Base route (get all products)
			productRoutes.GET("/", controllers.GetItems)

			// Create, Read, Update, Delete products
			productRoutes.GET("/:id", controllers.ItemHandler)
			productRoutes.POST("/", controllers.ItemHandler)      // (Admin Only)
			productRoutes.PUT("/:id", controllers.ItemHandler)    // (Admin Only)
			productRoutes.DELETE("/:id", controllers.ItemHandler) // (Admin Only)

			// Product variants
			productRoutes.GET("/:id/variants", controllers.BaseHandler)
			productRoutes.POST("/:id/variants", controllers.BaseHandler)     // (Admin Only)
			productRoutes.POST("/:id/variants/:id", controllers.BaseHandler) // (Admin Only)
			productRoutes.POST("/:id/variants/:id", controllers.BaseHandler) // (Admin Only)

			// Product inventory
			productRoutes.GET("/:id/stock", controllers.BaseHandler)
			productRoutes.PUT("/:id/stock", controllers.BaseHandler) // (Admin Only)

			// Product offers/deals
			productRoutes.GET("/offers", controllers.BaseHandler)
			productRoutes.PUT("/:id/offer", controllers.BaseHandler) // (Admin Only)

			// Rating and Reviews
			productRoutes.GET("/reviews", controllers.BaseHandler)
			productRoutes.GET("/:id/reviews", controllers.BaseHandler)
			productRoutes.POST("/:id/reviews", controllers.BaseHandler)
			productRoutes.PUT("/:id/reviews", controllers.BaseHandler)    // (Admin Only)
			productRoutes.DELETE("/:id/reviews", controllers.BaseHandler) // (Admin Only)

			// Search/Query products
			productRoutes.GET("/search", controllers.BaseHandler)
			productRoutes.GET("/search/category", controllers.BaseHandler)

			// Product Recommendations
			productRoutes.GET("/recommendations", controllers.BaseHandler)

			// Product Analyitics
			productRoutes.GET("/:id/analyitics", controllers.BaseHandler) // (Admin Only)

		}

	}

	return router
}
