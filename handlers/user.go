package handlers

import (
	"ecommerce-rbac/config"
	"ecommerce-rbac/models"
	"ecommerce-rbac/service"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignInRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"` // Minimum password length
}

func InitialUser() {
	log.Println("Initializing admin user...")

	var user models.User

	// Attempt to find the user with username "admin"
	err := config.DB.Where("username = ?", "admin").First(&user).Error

	// Handle unexpected errors during the query
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatalf("Error querying the database: %v", err)
	}

	// If the user already exists, log and return
	if err == nil {
		log.Printf("Admin user already exists: %s\n", user.Username)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash the password: %v", err)
	}

	// Create the initial admin user
	user = models.User{
		Username: "admin",
		Password: string(hashedPassword), // Store the hashed password
		Email:    "admin@gmail.com",
		Role:     "admin",
	}

	// Attempt to create the admin user
	if err := config.DB.Create(&user).Error; err != nil {
		log.Fatalf("Could not create admin user: %v", err)
	}

	log.Println("Admin user created successfully:", user.Username)
}

func CreateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Could not hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Create user model
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	// Save user to database
	if err := config.DB.Create(&user).Error; err != nil {
		log.Printf("Could not create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Generate JWT token for the new user
	token, err := service.GenerateJwtToken(user)
	if err != nil {
		log.Printf("Could not generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
		"token":   token, // Include JWT token in the response
	})
}

func SignIn(c *gin.Context) {
	var req SignInRequest

	// Bind JSON payload to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Fetch user from database by username or email
	err := config.DB.Where("username = ? OR email = ?", req.UsernameOrEmail, req.UsernameOrEmail).First(&user).Error
	if err != nil {
		log.Printf("User not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare provided password with stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Printf("Invalid password: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := service.GenerateJwtToken(user)
	if err != nil {
		log.Printf("Could not generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
		"token":   token, // Include JWT token in the response
	})
}
