package products

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"errors"
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

func VariantsHandler(c *gin.Context) {

	switch c.Request.Method {

	// Create variant
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
