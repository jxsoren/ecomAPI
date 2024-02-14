package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	// Load Env variabless
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env")
	}

	fmt.Println("Env Loaded Successfully!")

	// DB Environment Variables
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// DB Connection String
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	fmt.Println(dsn)

	// Connect to DB using DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// Try to get underlying MySQL db
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	// Try to ping DB to verify connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	// Creates Table (if it doesn't exist)
	db.AutoMigrate(&ProductItem{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	fmt.Println("Connected to DB")
}

type ProductItem struct {
	ID          string `gorm:"primaryKey;type:varchar(100)" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
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
	var productItem ProductItem

	// Find all records from product_items table
	results := db.Find(&productItem)
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

		var productItem ProductItem

		// Bind request body & check for errors
		if err := c.BindJSON(&productItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert new item into DB
		result := db.Create(&productItem)
		fmt.Println(productItem)

		// Throw 500 error if creation fails
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, productItem)

	case "PUT":
		var item Item

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
