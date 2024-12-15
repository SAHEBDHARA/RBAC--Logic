package models

import "gorm.io/gorm"


type Product struct {
	gorm.Model
	Name             string  `json:"name" gorm:"not null"`
	Description      string  `json:"description"`
	Price            float64 `json:"price" gorm:"not null"`
	Brand            string  `json:"brand"`
	ImageURL         string  `json:"image_url"`          // URL for the product image
	Category         string  `json:"category"`           // Category of the product
	Subcategory      string  `json:"subcategory"`        // Subcategory of the product
	ListedByUsername string  `json:"listed_by_username"` // Username of the user who listed the product
	ListedByRole     string  `json:"listed_by_role"`     // Role of the user who listed the product
}

