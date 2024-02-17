package main

import (
	"fmt"
	"log"

	"ecommerce_api/initializers"
	"ecommerce_api/routes"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

type Item struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

func main() {
	// Init Gin Router
	router := routes.SetupRouter()

	// Server Startup
	port := "8080"
	fmt.Printf("Server is running on port %s ...\n", port)
	log.Fatal(router.Run(":" + port))
}
