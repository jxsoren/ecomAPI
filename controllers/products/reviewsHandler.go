package products

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Review struct {
	ProductID    uint      `json:"product_id"`
	UserID       uint      `json:"user_id"`
	Rating       float32   `json:"rating"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ReviewDate   time.Time `json:"review_date"`
	IsVerified   bool      `json:"is_verified"`
	HelpfulCount int       `json:"helpful_count"`
}

type ReviewInput struct {
	ProductID    uint    `json:"product_id"`
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
	// Create instance of review input
	var reviewInput ReviewInput

	// Capture req body
	if err := c.BindJSON(&reviewInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Capture review input
	review := models.Review{
		ProductID:    reviewInput.ProductID,
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
	c.JSON(http.StatusOK, reviewInput)
}
