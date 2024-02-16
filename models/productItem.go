package models

import "gorm.io/gorm"

type ProductItem struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:varchar(100)" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}
