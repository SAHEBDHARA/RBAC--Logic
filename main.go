package main

import (
	"ecommerce-rbac/config"
	"ecommerce-rbac/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "service is up and running... 🚀🚀",
	})
}

func main() {
	// Connect to the database
	config.Connect()
	// handlers.InitialUser()
	port := config.AppConfig.Port
	if port == "" {
		port = "8080"
	}
	// Create a new Gin router
	router := gin.Default()

	// Define a simple route for testing
	router.GET("/", healthCheck)
	router.GET("/health", healthCheck)

	apiGroup := router.Group("/api/v1")
	routes.AuthRouters(apiGroup)
	routes.ProductRouter(apiGroup)

	// Start the server on port 8080
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	log.Println("Application started successfully on port 8080.")
}
