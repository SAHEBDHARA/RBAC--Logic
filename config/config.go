package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUsername      string
	DbPassword      string
	DbHost          string
	DbPort          string
	DbName          string
	Environment     string
	Port            string
	JwtSecretString string
}

var AppConfig Config

func init() {
	// Load .env file if it exists.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	AppConfig.DbUsername = os.Getenv("DB_USERNAME")
	AppConfig.DbPassword = os.Getenv("DB_PASSWORD")
	AppConfig.DbHost = os.Getenv("DB_HOST")
	AppConfig.DbPort = os.Getenv("DB_PORT")
	AppConfig.DbName = os.Getenv("DB_NAME")
	AppConfig.Environment = os.Getenv("ENVIRONMENT")
	AppConfig.Port = os.Getenv("PORT")
	AppConfig.JwtSecretString = os.Getenv("JWT_SECRET")

	if AppConfig.DbUsername == "" || AppConfig.DbPassword == "" || AppConfig.DbHost == "" ||
		AppConfig.DbPort == "" || AppConfig.DbName == "" {
		log.Fatal("One or more database configuration variables are missing.")
	}
}
