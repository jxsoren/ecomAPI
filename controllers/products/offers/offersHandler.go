package offers

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	helpers "ecommerce_api/utils"
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

func GetAllOffers(c *gin.Context) {
	// Query DB for all offers
	var offers []models.Offer
	if result := initializers.DB.First(&offers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Respond with offers
	c.JSON(http.StatusOK, offers)
}

func GetOffersForProduct(c *gin.Context) {
	// Validate that product exists
	id := c.Param("id")
	var product models.Product
	helpers.VerifyExistence(&product, id, c)

	// Query DB for all offers for product
	var offers []models.Offer
	if result := initializers.DB.First(&offers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Respond with all offers
	c.JSON(http.StatusOK, offers)
}

func G
