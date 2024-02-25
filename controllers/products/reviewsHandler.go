package products

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	UserID       uint    `json:"user_id"`
	Rating       float32 `json:"rating"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	IsVerified   bool    `json:"is_verified"`
	HelpfulCount int     `json:"helpful_count"`
}

func GetReviews(c *gin.Context) {
	// Create var to store results
	var reviews []models.Review

	// Query DB for reviews
	result := initializers.DB.Find(&reviews)

	// Handle Errors
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with reviews
	c.JSON(http.StatusOK, reviews)
}

func CreateReview(c *gin.Context) {
	// Caputre product ID from path param
	product_id := c.Param("id")

	// Convert product_id from string to int to uint
	productID, err := strconv.ParseInt(product_id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	} else if productID < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	uintProductID := uint(productID)

	// Verify product from path param exists
	var product models.Product
	productNotFoundMessage := fmt.Sprintf("Product with product ID of %s", product_id)
	if err := initializers.DB.First(&product, product_id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": productNotFoundMessage})
		return
	}

	// Capture req body
	var reviewInput ReviewInput
	if err := c.BindJSON(&reviewInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map review input to review model
	review := models.Review{
		ProductID:    uintProductID,
		UserID:       reviewInput.UserID,
		Rating:       reviewInput.Rating,
		Title:        reviewInput.Title,
		Content:      reviewInput.Content,
		ReviewDate:   time.Now(),
		IsVerified:   reviewInput.IsVerified,
		HelpfulCount: reviewInput.HelpfulCount,
	}

	// Create review in DB
	if result := initializers.DB.Create(&review); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": "error creating Review", "error": result.Error.Error()})
	}

	// Respond successfully with created review
	c.JSON(http.StatusOK, review)
}
