package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type item struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

var items = []item{
	{ID: "1", Description: "Item 1", Price: 100, Quantity: 10},
	{ID: "2", Description: "Item 2", Price: 200, Quantity: 20},
	{ID: "3", Description: "Item 3", Price: 300, Quantity: 30},
}

func main() {

	// Load Env variabless
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// DB Environment Variables
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// DB Connection String
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbHost, dbName)

	// Connect to DB using DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	// Verify connection to DB
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	log.Println("Connection to DB is sucessful")

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
	c.IndentedJSON(http.StatusOK, items)
}

func removeItem(s []item, index int) []item {
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

		var item item

		// Bind request body & check for errors
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		items = append(items, item)

		c.JSON(http.StatusOK, gin.H{
			"id":          item.ID,
			"description": item.Description,
			"price":       item.Price,
			"quantity":    item.Quantity,
		})

	case "PUT":
		var item item

		id := c.Param("id")

		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, item := range items {
			if item.ID == id {
				removeItem(items, i)

			}
		}

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
