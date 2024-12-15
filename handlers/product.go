package handlers

import (
	"ecommerce-rbac/config"
	"ecommerce-rbac/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req models.Product

	// Bind JSON payload to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve username and role from context
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Username not found"})
		return
	}

	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Role not found"})
		return
	}

	// Set ListedBy fields with user info
	req.ListedByUsername = username.(string)
	req.ListedByRole = role.(string)

	// Save product to database
	if err := config.DB.Create(&req).Error; err != nil {
		log.Printf("Could not create product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": req})
}