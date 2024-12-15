package config

import (
	"ecommerce-rbac/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", AppConfig.DbUsername, AppConfig.DbPassword, AppConfig.DbHost, AppConfig.DbPort, AppConfig.DbName)
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Database connected successfully.")
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.WebBuild{})
	// if err := DB.AutoMigrate(&models.Product{}); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }
	
	
}

