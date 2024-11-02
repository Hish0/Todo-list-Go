package main

import (
	"log"

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

	// Automatically migrate the User schema
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Setup Gin router
	router := gin.Default()

	// // Register API routes
	// router.POST("/register", func(c *gin.Context) {
	// 	handlers.Register(c, db)
	// })
	// router.POST("/login", func(c *gin.Context) {
	// 	handlers.Login(c, db)
	// })

	// Setup routes using the routes package
	routes.SetupRoutes(router, db)

	// Run the Gin server
	router.Run(":8080")
}
