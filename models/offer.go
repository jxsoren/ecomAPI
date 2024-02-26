package models

import (
	"time"

	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	ProductID      uint      `gorm:"foreignKey:ProductID" json:"product_id"`
	Title          string    `gorm:"type:varchar(255)" json:"name"`
	Description    string    `gorm:"type:text" json:"description"`
	DiscountRate   float64   `gorm:"type:decimal(3, 2)" json:"discount_rate"`
	DiscountAmount float64   `gorm:"type:decimal(3, 2)" json:"discount_amount"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	PromoCode      string    `gorm:"type:varchar(255)" json:"promo_code"`
	MinPurchase    float64   `gorm:"type:decimal(3, 2)" json:"min_purchase"`
	IsActive       bool      `json:"is_active"`
}
