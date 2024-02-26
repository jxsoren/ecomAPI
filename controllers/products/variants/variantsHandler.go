package variants

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VariantInput struct {
	ProductID             uint    `json:"product_id"`
	Size                  string  `json:"size"`
	Color                 string  `json:"color"`
	Material              string  `json:"material"`
	Weight                float32 `json:"weight"`
	Dimensions            string  `json:"dimensions"`
	SKU                   string  `json:"sku"`
	Price                 float64 `json:"price"`
	StockQuantity         int     `json:"stock_quantity"`
	ImageURL              string  `json:"image_url"`
	AdditionalDescription string  `json:"additiona_description"`
	SalePrice             float32 `json:"sale_price"`
	AvailabilityStatus    string  `json:"availability_status"`
	ShippingDetails       string  `json:"shipping_details"`
}

type VariantUpdateInput struct {
	Size                  *string  `json:"size"`
	Color                 *string  `json:"color"`
	Material              *string  `json:"material"`
	Weight                *float32 `json:"weight"`
	Dimensions            *string  `json:"dimensions"`
	SKU                   *string  `json:"sku"`
	Price                 *float64 `json:"price"`
	StockQuantity         *int     `json:"stock_quantity"`
	ImageURL              *string  `json:"image_url"`
	AdditionalDescription *string  `json:"additiona_description"`
	SalePrice             *float32 `json:"sale_price"`
	AvailabilityStatus    *string  `json:"availability_status"`
	ShippingDetails       *string  `json:"shipping_details"`
}

func VariantsHandler(c *gin.Context) {

	switch c.Request.Method {

	case "GET":
		// Get IDs from path params
		product_id := c.Param("id")
		variant_id := c.Param("variant_id")

		// Create instances of product and vairant models
		var product models.Product
		var variant models.Variant

		// Query DB for product id
		result := initializers.DB.First(&product, product_id)

		// Check for not record not found + handle catch all error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Get variant from DB
		if result := initializers.DB.First(&variant, variant_id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error, "error_message": "Variant not found"})
			return
		}

		// Return variant
		c.JSON(http.StatusOK, variant)

	case "POST":
		// Create instance of variant input
		var input VariantInput

		// Handle errors from client request
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Map request body to model
		variant := models.Variant{
			ProductID:             input.ProductID,
			Size:                  input.Size,
			Color:                 input.Color,
			Material:              input.Color,
			Weight:                input.Weight,
			Dimensions:            input.Color,
			SKU:                   input.Color,
			Price:                 input.Price,
			StockQuantity:         input.StockQuantity,
			ImageURL:              input.ImageURL,
			AdditionalDescription: input.AdditionalDescription,
			SalePrice:             input.SalePrice,
			AvailabilityStatus:    input.AvailabilityStatus,
			ShippingDetails:       input.ShippingDetails,
		}

		// Add variant to DB & handle errors
		if result := initializers.DB.Create(&variant); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}

		// Respond with successful message
		c.JSON(http.StatusOK, variant)

	case "PUT":
		// Get IDs from path params
		product_id := c.Param("id")
		variant_id := c.Param("variant_id")

		// Verify product exists
		var product models.Product
		productResult := initializers.DB.First(&product, product_id)
		productNotFoundMessage := fmt.Sprintf("No product with ID of %s found.", product_id)
		if errors.Is(productResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": productNotFoundMessage})
			return
		} else if productResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": productResult.Error.Error()})
			return
		}

		// Verify variant exists
		var variant models.Variant
		variantResult := initializers.DB.First(&variant, variant_id)
		reviewNotFoundMessage := fmt.Sprintf("No review with ID of %s found.", variant_id)
		if errors.Is(variantResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": reviewNotFoundMessage})
			return
		} else if variantResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": variantResult.Error.Error()})
			return
		}

		// Capture request body
		var variantUpdateInput VariantUpdateInput
		if err := c.BindJSON(&variantUpdateInput); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}

		if variantUpdateInput.StockQuantity != nil {
			variant.StockQuantity = *variantUpdateInput.StockQuantity
		}
		if variantUpdateInput.SKU != nil {
			variant.SKU = *variantUpdateInput.SKU
		}
		if variantUpdateInput.ImageURL != nil {
			variant.ImageURL = *variantUpdateInput.ImageURL
		}
		if variantUpdateInput.Weight != nil {
			variant.Weight = *variantUpdateInput.Weight
		}
		if variantUpdateInput.Dimensions != nil {
			variant.Dimensions = *variantUpdateInput.Dimensions
		}
		if variantUpdateInput.Color != nil {
			variant.Color = *variantUpdateInput.Color
		}
		if variantUpdateInput.Size != nil {
			variant.Size = *variantUpdateInput.Size
		}

		// Update variant in DB
		if result := initializers.DB.Save(&variant); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}

		// Respond sucessfully
		c.JSON(http.StatusOK, variant)
	case "DELETE":
		// Create instance of variant model
		var variant models.Variant
		var product models.Product

		// Get IDs from path params
		productID := c.Param("id")
		variantID := c.Param("variant_id")

		// Check that both product and variant records exists & handle according errors
		if productResult := initializers.DB.First(&product, productID); errors.Is(productResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No product found"})
			return
		}

		if variantResult := initializers.DB.Find(&variant, variantID); errors.Is(variantResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No variant found"})
			return
		}

		// Remove variant record from DB & handle errors
		if result := initializers.DB.Delete(&variant, variantID); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Respond with successful status
		c.JSON(http.StatusAccepted, gin.H{"error": "Variant deleted successfully"})

	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
