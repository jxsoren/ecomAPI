package main

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"log"
)

func init() {
	// Load Env variables
	initializers.LoadEnvVars()
	initializers.ConnectToDB()

}

func main() {

	// Creates Table (if it doesn't exist)
	err := initializers.DB.AutoMigrate(&models.ProductItem{})

	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
}
