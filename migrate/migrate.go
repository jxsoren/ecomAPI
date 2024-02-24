package main

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"log"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {

	if productsErr := initializers.DB.AutoMigrate(&models.Product{}); productsErr != nil {
		log.Fatalf("Error during migration: %v", productsErr)
	}

	if variantsErr := initializers.DB.AutoMigrate(&models.Variant{}); variantsErr != nil {
		log.Fatalf("Error during Variant Table migration. Error: %v", variantsErr)
	}

	if reviewsErr := initializers.DB.AutoMigrate(&models.Review{}); reviewsErr != nil {
		log.Fatalf("Error during Reviews Table Migration. Error: %v", reviewsErr)
	}
}
