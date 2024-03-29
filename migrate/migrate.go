package main

import (
	"ecommerce_api/initializers"
	"ecommerce_api/models"
	"fmt"
	"log"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {

	if productsErr := initializers.DB.AutoMigrate(&models.Product{}); productsErr != nil {
		log.Fatalf("Error during migration: %v", productsErr)
		return
	}

	fmt.Println("Procucts migrated successfully!")

	if variantsErr := initializers.DB.AutoMigrate(&models.Variant{}); variantsErr != nil {
		log.Fatalf("Error during Variant Table migration. Error: %v", variantsErr)
		return
	}

	fmt.Println("Variants migrated successfully!")

	if reviewsErr := initializers.DB.AutoMigrate(&models.Review{}); reviewsErr != nil {
		log.Fatalf("Error during Reviews Table Migration. Error: %v", reviewsErr)
		return
	}

	fmt.Println("Reviews migrated successfully!")

	if offersErr := initializers.DB.AutoMigrate(&models.Offer{}); offersErr != nil {
		log.Fatalf("Error during Reviews Table Migration. Error: %v", offersErr)
		return
	}

	fmt.Println("Offers migrated successfully!")

}
