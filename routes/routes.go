package routes

import (
	"ecommerce_api/controllers"
	"ecommerce_api/controllers/products"
	"ecommerce_api/controllers/products/offers"
	"ecommerce_api/controllers/products/reviews"
	"ecommerce_api/controllers/products/stock"
	"ecommerce_api/controllers/products/variants"

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
			productRoutes.GET("/", products.GetProducts)

			// Create, Read, Update, Delete products
			productRoutes.GET("/:id", products.ProductsHandler)
			productRoutes.POST("/", products.ProductsHandler)      // ! (Admin Only)
			productRoutes.PUT("/:id", products.ProductsHandler)    // ! (Admin Only)
			productRoutes.DELETE("/:id", products.ProductsHandler) // ! (Admin Only)

			// Product variants
			productRoutes.GET("/:id/variants/:variant_id", variants.VariantsHandler)
			productRoutes.POST("/:id/variants", variants.VariantsHandler)               // ! (Admin Only)
			productRoutes.PUT("/:id/variants/:variant_id", variants.VariantsHandler)    // ! (Admin Only)
			productRoutes.DELETE("/:id/variants/:variant_id", variants.VariantsHandler) // ! (Admin Only)

			// Variant inventory
			productRoutes.GET("/:id/variants/:variant_id/stock", stock.GetVariantStock)
			productRoutes.PUT("/:id/variants/:variant_id/stock", stock.UpdateVariantStock) // ! (Admin Only)

			// Product inventory
			productRoutes.GET("/:id/stock", stock.GetStock)
			productRoutes.PUT("/:id/stock", stock.UpdateStock) // ! (Admin Only)

			// TODO: create handlers for routes
			// Product offers/deals
			productRoutes.GET("/offers", offers.GetAllOffers)
			productRoutes.POST("/offers", offers.CreateOffer)
			productRoutes.GET("/:id/offer")
			productRoutes.GET("/:id/offers", offers.GetOffersForProduct)
			productRoutes.PUT("/:id/offer")    // ! (Admin Only)
			productRoutes.DELETE("/:id/offer") // ! (Admin Only)

			// TODO: refactor & review
			// Rating and Reviews
			productRoutes.GET("/reviews", reviews.GetAllReviews)
			productRoutes.GET("/:id/reviews", reviews.GetAllReivewsForProduct)
			productRoutes.GET("/:id/reviews/:review_id", reviews.GetReview)
			productRoutes.POST("/:id/reviews", reviews.CreateReview)
			productRoutes.PUT("/:id/reviews/:review_id", reviews.UpdateReview)
			productRoutes.DELETE("/:id/reviews/:review_id", reviews.DeleteReveiew) // ! (Admin Only)

			// TODO: create handlers for routes
			// Search/Query products
			productRoutes.GET("/search")
			productRoutes.GET("/search/category")

			// TODO: create handlers for routes
			// Product Recommendations
			productRoutes.GET("/recommendations")

			// TODO: create handlers for routes
			// Product Analyitics
			productRoutes.GET("/:id/analyitics") // ! (Admin Only)
		}

		userRoutes := v1Routes.Group("/users")
		{
			// TODO: create handlers for routes
			// User registration
			userRoutes.POST("/signup")

			// TODO: create handlers for routes
			// User authentication
			userRoutes.POST("/authentication")
			userRoutes.POST("logout")

			// TODO: create handlers for routes
			// Profile management
			userRoutes.GET("/profile")
			userRoutes.PUT("/profile")
			userRoutes.DELETE("/profile")

			// TODO: create handlers for routes
			// Password management
			userRoutes.POST("/forgot-password")
			userRoutes.POST("/reset-password")

			// TODO: create handlers for routes
			// User administration
			userRoutes.GET("/users")     // ! (Admin Only)
			userRoutes.GET("/users/:id") // ! (Admin Only)
			userRoutes.PUT("deactivate") // ! (Admin Only)
			userRoutes.PUT("activate")   // ! (Admin Only)
		}

		cartRoutes := v1Routes.Group("/cart")
		{
			// TODO: create handlers for routes
			// View cart
			cartRoutes.GET("/")

			// TODO: create handlers for routes
			// Get cart item count
			cartRoutes.GET("/count")

			// TODO: create handlers for routes
			// Add, update and delete cart items
			cartRoutes.POST("/add")
			cartRoutes.PUT("/update")
			cartRoutes.DELETE("/remove")

			// TODO: create handlers for routes
			// Empty cart
			cartRoutes.POST("/cart/empty")

			// TODO: create handlers for routes
			// Checkout
			cartRoutes.POST("cart/checkout")
		}

		orderRoutes := v1Routes.Group("/orders")
		{
			// TODO: create handlers for routes
			// Create order
			orderRoutes.POST("/")

			// TODO: create handlers for routes
			// Get all orders
			orderRoutes.GET("/") // ! (Admin Only)

			// TODO: create handlers for routes
			// List all orders for user
			orderRoutes.GET("/user/:id")

			// TODO: create handlers for routes
			// Update order status
			orderRoutes.PUT("/:id/status")

			// TODO: create handlers for routes
			// Cancel an order
			orderRoutes.PUT("/:id/cancel")

			// TODO: create handlers for routes
			// Return order
			orderRoutes.POST("/:id/return")
		}

	}

	return router
}
