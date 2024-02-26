package stock

import (
	"ecommerce_api/initializers"
	helpers "ecommerce_api/utils"
	"net/http"

	"ecommerce_api/models"

	"github.com/gin-gonic/gin"
)

type StockUpdateInput struct {
	StockQuantity *int `json:"stock_quantity"`
}

func GetStock(c *gin.Context) {
	// Create instance of Product
	var product models.Product

	// Get ID from path param
	id := c.Param("id")

	// Query DB for product ID, save result and handle errors
	if result := initializers.DB.Find(&product, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond back with quantity and successful status
	c.JSON(http.StatusOK, gin.H{"stock_quantity": product.StockQuantity})
}

func UpdateStock(c *gin.Context) {
	// Create instance of Product
	var product models.Product

	// Get ID from path parm
	id := c.Param("id")

	// Create instance of StockUpdateInput
	var stockUpdateInput StockUpdateInput

	// Bind req body to product
	if err := c.BindJSON(&stockUpdateInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verify that product exists
	if err := initializers.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error})
		return
	}

	// Validate inputs
	if stockUpdateInput.StockQuantity != nil {
		product.StockQuantity = *stockUpdateInput.StockQuantity
	}

	// Make DB call to update stock quantity field
	if result := initializers.DB.Save(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond successfully
	c.JSON(http.StatusOK, gin.H{"status": "successful", "stock_quantity": product.StockQuantity})
}

func GetVariantStock(c *gin.Context) {
	// Get IDs from path params
	product_id := c.Param("id")
	variant_id := c.Param("variant_id")

	// Verify that product exists in DB
	var product models.Product
	helpers.VerifyExistence(&product, product_id, c)

	// Verity that variant exists in DB
	var variant models.Variant
	helpers.VerifyExistence(&variant, variant_id, c)

	// Respond with variant stock
	c.JSON(http.StatusOK, gin.H{"stock_quantity": variant.StockQuantity})
}

func UpdateVariantStock(c *gin.Context) {
	// Get IDs from path params
	product_id := c.Param("id")
	variant_id := c.Param("variant_id")

	// Verify that product exists in DB
	var product models.Product
	helpers.VerifyExistence(&product, product_id, c)

	// Verity that variant exists in DB
	var variant models.Variant
	helpers.VerifyExistence(&variant, variant_id, c)

	// Capture req body
	var input StockUpdateInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Assign new input value to instance of model
	if input.StockQuantity != nil {
		variant.StockQuantity = *input.StockQuantity
	}

	// Update DB with new stock value
	if result := initializers.DB.Save(&variant); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond successfully with new stock value
	c.JSON(http.StatusOK, gin.H{"status": "successful", "stock_quantity": variant.StockQuantity})
}
