package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin Router
	router := gin.Default()

	// Routes
	router.GET("/items", getItems)
	router.GET("/items/:id", itemHandler)
	router.POST("/items", itemHandler)
	router.PUT("/items/:id", itemHandler)
	router.DELETE("/items/:id", itemHandler)
	router.GET("/", baseHandler)

	// Server Startup
	port := "8080"
	fmt.Printf("Server is running on port %s ...\n", port)
	log.Fatal(router.Run(":" + port))
}

func getItems(c *gin.Context) {
	c.String(http.StatusOK, "Retrieving all items")
}

func itemHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Param("id")
		c.String(http.StatusOK, "Retrieving item with ID %s", id)
	case "POST":
		c.String(http.StatusOK, "Creating an item")
	case "PUT":
		id := c.Param("id")
		c.String(http.StatusOK, "Updating item with ID %s", id)
	case "DELETE":
		id := c.Param("id")
		c.String(http.StatusOK, "Deleting item with ID %s", id)
	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func baseHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello!!! 🐹 🐹 🐹 \n")
}
