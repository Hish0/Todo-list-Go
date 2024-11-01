package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv" // Add this line
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDatabase establishes a connection to the PostgreSQL database
func ConnectToDatabase() (*gorm.DB, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Database connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
