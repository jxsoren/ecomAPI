package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce_api/initializers"
	"ecommerce_api/models"

	"github.com/gin-gonic/gin"
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

var items = []Item{
	{ID: "1", Description: "Item 1", Price: 100, Quantity: 10},
	{ID: "2", Description: "Item 2", Price: 200, Quantity: 20},
	{ID: "3", Description: "Item 3", Price: 300, Quantity: 30},
}

func main() {
	// Init Gin Router
	router := gin.Default()

	// Routes
	router.GET("/items", getItems)

	// Item ID Routes
	router.GET("/items/:id", itemHandler)
	router.POST("/items", itemHandler)
	router.PUT("/items/:id", itemHandler)
	router.DELETE("/items/:id", itemHandler)

	// Base Route
	router.GET("/", baseHandler)

	// Server Startup
	port := "8080"
	fmt.Printf("Server is running on port %s ...\n", port)
	log.Fatal(router.Run(":" + port))
}

func getItems(c *gin.Context) {
	var productItem models.ProductItem

	// Find all records from product_items table
	results := initializers.DB.Find(&productItem)
	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
	}

	// Return items in pretty printed JSON -- may need to return compact JSON in prod
	c.IndentedJSON(http.StatusOK, productItem)
}

func removeItem(s []Item, index int) []Item {
	return append(s[:index], s[index+1:]...)
}

func itemHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Param("id")
		c.String(http.StatusOK, "Retrieving item with ID %s \n", id)

		for _, item := range items {
			if item.ID == id {
				c.IndentedJSON(http.StatusOK, item)
				return
			}
		}
	case "POST":

		var productItem models.ProductItem

		// Bind request body & check for errors
		if err := c.BindJSON(&productItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert new item into DB
		result := initializers.DB.Create(&productItem)
		fmt.Println(productItem)

		// Throw 500 error if creation fails
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, productItem)

	case "PUT":
		var productItem models.ProductItem

		id := c.Param("id")

		// Handle error
		if err := c.BindJSON(&productItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// result := initializers.DB.Update()

		c.String(http.StatusOK, "Updating item with ID %s", id)
	case "DELETE":
		id := c.Param("id")

		for i, item := range items {
			if item.ID == id {
				removeItem(items, i)
				c.JSON(http.StatusNoContent, gin.H{})
				return
			}
		}

	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func baseHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello!!! 🐹 🐹 🐹 \n")
}
