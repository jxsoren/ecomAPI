package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Testing 123")

	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/", itemHandler)
	http.HandleFunc("/", baseHandler)

	port := "8080"
	newPort := ":" + port
	fmt.Printf("Server is running on port %s ...\n", port)

	log.Fatal(http.ListenAndServe(newPort, nil))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Retrieving all items")
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Handle GET for specific item
		fmt.Fprintf(w, "Retrieving an item")
	case "POST":
		// Handle POST to create an item
		fmt.Fprintf(w, "Creating an item")
	case "PUT":
		// Handle PUT to update an item
		fmt.Fprintf(w, "Updating an item")
	case "DELETE":
		// Handle DELETE to delete an item
		fmt.Fprintf(w, "Deleting an item")
	default:
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!! 🐹 🐹 🐹 \n")
}
