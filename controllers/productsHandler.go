package controllers

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var productItem models.Product

	// Find all records from product_items table
	results := initializers.DB.Find(&productItem)
	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
	}

	// Return items in pretty printed JSON -- may need to return compact JSON in prod
	c.IndentedJSON(http.StatusOK, productItem)
}

func ProductsHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		// Create ref to Product model
		var productItem models.Product

		// Get product id from path param
		id := c.Param("id")

		// Query DB for product id
		result := initializers.DB.First(&models.Product{}, id)

		// Check for errors
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}

		// Return product item
		c.JSON(http.StatusOK, productItem)

	case "POST":

		var productItem models.Product

		// Bind request body & check for errors
		if err := c.BindJSON(&productItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert new item into DB
		result := initializers.DB.Create(&productItem)
		fmt.Println(productItem)

		// Throw 500 error if creation fails
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, productItem)

	case "PUT":
		var productItem models.Product

		id := c.Param("id")

		// Handle error
		if err := c.BindJSON(&productItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// result := initializers.DB.Update()

		c.String(http.StatusOK, "Updating item with ID %s", id)
	case "DELETE":
		// id := c.Param("id")

	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
