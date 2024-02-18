package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255)" json:"name"`
	Description   string    `gorm:"type:text" json:"description"`
	Price         float64   `gorm:"type:decimal(10,2)" json:"price"`
	CategoryID    int       `gorm:"index" json:"category_id"`
	StockQuantity int       `json:"stock_quantity"`
	SKU           string    `gorm:"type:varchar(100)" json:"sku"`
	ImageURL      string    `gorm:"type:varchar(255)" json:"image_url"`
	AddedDate     time.Time `json:"added_date"`
	UpdatedDate   time.Time `json:"updated_date"`
	IsActive      bool      `json:"is_active"`
	Weight        float32   `gorm:"type:decimal(5,2)" json:"weight"`
	Dimensions    string    `gorm:"type:varchar(100)" json:"dimensions"`
	Color         string    `gorm:"type:varchar(50)" json:"color"`
	Size          string    `gorm:"type:varchar(50)" json:"size"`
	Rating        float32   `gorm:"type:decimal(3,2)" json:"rating"`
}
