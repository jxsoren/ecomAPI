package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ProductID    uint      `gorm:"index" json:"product_id"` // Foreign key for Product
	UserID       uint      `gorm:"index" json:"user_id"`
	Rating       float32   `gorm:"type:decimal(3,2)" json:"rating"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	Content      string    `gorm:"type:text" json:"content"`
	ReviewDate   time.Time `json:"review_date"`
	IsVerified   bool      `json:"is_verified"`
	HelpfulCount int       `json:"helpful_count"`
}
