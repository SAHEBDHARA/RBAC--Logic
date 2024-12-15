package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Role     string `json:"role" gorm:"not null"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"` // Assuming Role is defined in constants package
}


type UserClaims struct {
    Username string
    Role     string
}