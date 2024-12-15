package handlers

import (
	"ecommerce-rbac/config"
	"ecommerce-rbac/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateWebBuild handles creating a new web build
func CreateWebBuild(c *gin.Context) {
	var req models.WebBuild

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

	// Save web build to database
	if err := config.DB.Create(&req).Error; err != nil {
		log.Printf("Could not create web build: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create web build"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Web build created successfully", "web_build": req})
}
