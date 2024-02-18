package controllers

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"fmt"
	"net/http"
	"time"

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

type ProductInput struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryID    int     `json:"category_id"`
	StockQuantity int     `json:"stock_quantity"`
	SKU           string  `json:"sku"`
	ImageURL      string  `json:"image_url"`
	Weight        float32 `json:"weight"`
	Dimensions    string  `json:"dimensions"`
	Color         string  `json:"color"`
	Size          string  `json:"size"`
}

func ProductsHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		// Create ref to Product model
		var productItem models.Product

		// Get product id from path param
		id := c.Param("id")

		// Query DB for product id
		result := initializers.DB.First(&productItem, id)

		// Check for errors
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}

		// Return product item
		c.JSON(http.StatusOK, productItem)

	case "POST":
		// Create instance of product input
		var productInput ProductInput

		// Bind req body to productInput & handle errors
		if err := c.BindJSON(&productInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Map productItem values to instance of Product model
		productItem := models.Product{
			Name:          productInput.Name,
			Description:   productInput.Description,
			Price:         productInput.Price,
			CategoryID:    productInput.CategoryID,
			StockQuantity: productInput.StockQuantity,
			SKU:           productInput.SKU,
			ImageURL:      productInput.ImageURL,
			AddedDate:     time.Now(),
			UpdatedDate:   time.Now(),
			IsActive:      true,
			Weight:        productInput.Weight,
			Dimensions:    productInput.Dimensions,
			Color:         productInput.Color,
			Size:          productInput.Size,
			Rating:        000.00,
		}

		// Insert productItem into DB
		result := initializers.DB.Create(&productItem)
		fmt.Println(productItem)

		// Handle error
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Return productItem back to client
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
