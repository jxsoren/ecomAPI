package products

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context) {
	var productItem models.Product

	// Find all records from product_items table
	result := initializers.DB.Find(&productItem)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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

// All types are set as pointers to handle optional fields
type ProductUpdateInput struct {
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	CategoryID    *int     `json:"category_id"`
	StockQuantity *int     `json:"stock_quantity"`
	SKU           *string  `json:"sku"`
	ImageURL      *string  `json:"image_url"`
	Weight        *float32 `json:"weight"`
	Dimensions    *string  `json:"dimensions"`
	Color         *string  `json:"color"`
	Size          *string  `json:"size"`
	Rating        *float32 `json:"rating"`
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

		// Check for not record not found + handle catch all error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
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
		// Create inst of ProductInput struct (all values are pointers to types)
		var input ProductUpdateInput

		// Get product ID from path param
		id := c.Param("id")

		// Handle input binding errors
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Find product using ID
		var product models.Product

		// Handle errors if product is unfound
		notFoundMessage := fmt.Sprintf("Product not found with ID of %s", id)
		if err := initializers.DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
			return
		}

		// Validate inputs
		if input.Name != nil {
			product.Name = *input.Name
		}
		if input.Description != nil {
			product.Description = *input.Description
		}
		if input.Price != nil {
			product.Price = *input.Price
		}
		if input.CategoryID != nil {
			product.CategoryID = *input.CategoryID
		}
		if input.StockQuantity != nil {
			product.StockQuantity = *input.StockQuantity
		}
		if input.SKU != nil {
			product.SKU = *input.SKU
		}
		if input.ImageURL != nil {
			product.ImageURL = *input.ImageURL
		}
		if input.Weight != nil {
			product.Weight = *input.Weight
		}
		if input.Dimensions != nil {
			product.Dimensions = *input.Dimensions
		}
		if input.Color != nil {
			product.Color = *input.Color
		}
		if input.Size != nil {
			product.Size = *input.Size
		}
		if input.Rating != nil {
			product.Rating = *input.Rating
		}

		// Update product
		result := initializers.DB.Save(&product)

		// Handle handle update errors
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Respond with updated product
		c.JSON(http.StatusOK, product)
	case "DELETE":
		// Create inst of product
		var product models.Product

		// Get ID from path param
		id := c.Param("id")

		// Check for product & handle not found
		notFoundMessage := fmt.Sprintf("Could not delete. No products found with ID of %s", id)
		if result := initializers.DB.First(&product, id); result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Delete product from DB
		initializers.DB.Delete(&product)

		// Respond with 200
		c.JSON(http.StatusOK, gin.H{"status": "Product successfully deleted!"})

	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
