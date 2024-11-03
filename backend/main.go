package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo-list-go/backend/database" // Import database package
	"todo-list-go/backend/models"
	"todo-list-go/backend/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to the database using the function in database.go
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Automatically migrate the User and Task schema
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Setup Gin router
	router := gin.Default()

	// Enable CORS to let frontend comunicate with backend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow requests from your React app
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Setup routes using the routes package
	routes.SetupRoutes(router, db)

	// Run the Gin server
	router.Run(":8080")
}
