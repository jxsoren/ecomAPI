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

func GetReview(c *gin.Context) {
	// Capture IDs from path params
	product_id := c.Param("id")
	review_id := c.Param("review_id")

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

	// Retrieve review from DB
	var review models.Review
	reviewResult := initializers.DB.First(&review, review_id)
	reviewNotFoundMessage := fmt.Sprintf("No review with ID of %s found.", review_id)
	if errors.Is(reviewResult.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": reviewNotFoundMessage})
		return
	} else if reviewResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": reviewResult.Error.Error()})
		return
	}

	// Respond sucessfully with review
	c.JSON(http.StatusOK, review)
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
	result := initializers.DB.First(&product, product_id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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
