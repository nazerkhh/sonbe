// cmd/apiserver/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nazerkh09/sonbe/internal/app/apiserver/database"
	"github.com/Nazerkh09/sonbe/internal/app/apiserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Run the server on port 8080
	port := 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost%s\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
