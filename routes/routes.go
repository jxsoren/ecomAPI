package routes

import (
	"ecommerce_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Root Route
	router.GET("/", controllers.BaseHandler)

	apiRoutes := router.Group("/api")
	v1Routes := apiRoutes.Group("/v1")
	{

		// Product Routes
		productRoutes := v1Routes.Group("/products")
		{
			// Base route (get all products)
			productRoutes.GET("/", controllers.GetProducts)

			// Create, Read, Update, Delete products
			productRoutes.GET("/:id", controllers.ProductsHandler)
			productRoutes.POST("/", controllers.ProductsHandler)      // (Admin Only)
			productRoutes.PUT("/:id", controllers.ProductsHandler)    // (Admin Only)
			productRoutes.DELETE("/:id", controllers.ProductsHandler) // (Admin Only)

			// Product variants
			productRoutes.GET("/:id/variants")
			productRoutes.POST("/:id/variants")       // (Admin Only)
			productRoutes.POST("/:id/variants/:id")   // (Admin Only)
			productRoutes.DELETE("/:id/variants/:id") // (Admin Only)

			// Product inventory
			productRoutes.GET("/:id/stock")
			productRoutes.PUT("/:id/stock") // (Admin Only)

			// Product offers/deals
			productRoutes.GET("/offers")
			productRoutes.PUT("/:id/offer") // (Admin Only)

			// Rating and Reviews
			productRoutes.GET("/reviews")
			productRoutes.GET("/:id/reviews")
			productRoutes.POST("/:id/reviews")
			productRoutes.PUT("/:id/reviews")    // (Admin Only)
			productRoutes.DELETE("/:id/reviews") // (Admin Only)

			// Search/Query products
			productRoutes.GET("/search")
			productRoutes.GET("/search/category")

			// Product Recommendations
			productRoutes.GET("/recommendations")

			// Product Analyitics
			productRoutes.GET("/:id/analyitics") // (Admin Only)
		}

		userRoutes := v1Routes.Group("/users")
		{
			// User registration
			userRoutes.POST("/signup")

			// User authentication
			userRoutes.POST("/authentication")
			userRoutes.POST("logout")

			// Profile management
			userRoutes.GET("/profile")
			userRoutes.PUT("/profile")
			userRoutes.DELETE("/profile")

			// Password management
			userRoutes.POST("/forgot-password")
			userRoutes.POST("/reset-password")

			// User administration
			userRoutes.GET("/users")     // (Admin Only)
			userRoutes.GET("/users/:id") // (Admin Only)
			userRoutes.PUT("deactivate") // (Admin Only)
			userRoutes.PUT("activate")   // (Admin Only)
		}

		cartRoutes := v1Routes.Group("/cart")
		{
			// View cart
			cartRoutes.GET("/")

			// Get cart item count
			cartRoutes.GET("/count")

			// Add, update and delete cart items
			cartRoutes.POST("/add")
			cartRoutes.PUT("/update")
			cartRoutes.DELETE("/remove")

			// Empty cart
			cartRoutes.POST("/cart/empty")

			// Checkout
			cartRoutes.POST("cart/checkout")
		}

		orderRoutes := v1Routes.Group("/orders")
		{
			// Create order
			orderRoutes.POST("/")

			// Get all orders
			orderRoutes.GET("/") // (Admin Only)

			// List all orders for user
			orderRoutes.GET("/user/:id")

			// Update order status
			orderRoutes.PUT("/:id/status")

			// Cancel an order
			orderRoutes.PUT("/:id/cancel")

			// Return order
			orderRoutes.POST("/:id/return")
		}

	}

	return router
}
