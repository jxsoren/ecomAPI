package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ProductID             uint    `gorm:"foreignKey:ProductID" json:"product_id"`
	Size                  string  `gorm:"type:varchar(50)" json:"size"`
	Color                 string  `gorm:"type:varchar(50)" json:"color"`
	Material              string  `gorm:"type:varchar(50)" json:"material"`
	Weight                float32 `gorm:"type:decimal(5,2)" json:"weight"`
	Dimensions            string  `gorm:"type:varchar(100)" json:"dimensions"`
	SKU                   string  `gorm:"type:varchar(100)" json:"sku"`
	Price                 float64 `gorm:"type:decimal(10,2)" json:"price"`
	StockQuantity         int     `gorm:"type:int" json:"stock_quantity"`
	ImageURL              string  `gorm:"type:varchar(255)" json:"image_url"`
	AdditionalDescription string  `gorm:"type:text" json:"additiona_description"`
	SalePrice             float32 `gorm:"type:decimal(10,2)" json:"sale_price"`
	AvailabilityStatus    string  `gorm:"type:varchar(10)" json:"availability_status"`
	ShippingDetails       string  `gorm:"type:text" json:"shipping_details"`
	CategoryID            int     `gorm:"index" json:"category_id"`
}
