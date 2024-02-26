package offers

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OfferInput struct {
	Title          string    `json:"name"`
	Description    string    `json:"description"`
	DiscountRate   float64   `json:"discount_rate"`
	DiscountAmount float64   `json:"discount_amount"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	PromoCode      string    `json:"promo_code"`
	MinPurchase    float64   `json:"min_purchase"`
	IsActive       bool      `json:"is_active"`
}

type OfferUpdateInput struct {
	Title          string    `json:"name"`
	Description    string    `json:"description"`
	DiscountRate   float64   `json:"discount_rate"`
	DiscountAmount float64   `json:"discount_amount"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	PromoCode      string    `json:"promo_code"`
	MinPurchase    float64   `json:"min_purchase"`
	IsActive       bool      `json:"is_active"`
}

func CreateOffer(c *gin.Context) {
	// Capture request body
	var input OfferInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Bind input to inst of offer model
	newOffer := models.Offer{
		Title:          input.Title,
		Description:    input.Description,
		DiscountAmount: input.DiscountAmount,
		StartDate:      input.StartDate,
		EndDate:        input.EndDate,
		PromoCode:      input.PromoCode,
		MinPurchase:    input.MinPurchase,
		IsActive:       input.IsActive,
	}

	// Create offer in DB
	if result := initializers.DB.Create(&newOffer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Respond successfully w/ created offer
	c.JSON(http.StatusCreated, newOffer)
}
