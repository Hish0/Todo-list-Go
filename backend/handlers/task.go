package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"todo-list-go/backend/models" // Import models package
)

// CreateTask handler for creating a new task
func CreateTask(c *gin.Context, db *gorm.DB) {
	// Retrieve the user ID from the JWT token
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	var task models.Task // Use the Task struct from models package
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the task title and description (ensure they are not empty)
	if task.Title == "" || task.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Description are required"})
		return
	}

	// Set the UserID for the task
	task.UserID = userID

	// Create the task in the database
	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Return the created task
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": task})
}

// GetTasks handler for retrieving tasks for a user
func GetTasks(c *gin.Context, db *gorm.DB) {
	// Retrieve the user ID from the JWT token
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Retrieve all tasks for the user from the database
	var tasks []models.Task // Use the Task struct from models package
	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	// Return the tasks
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// UpdateTask handler for updating a task
func UpdateTask(c *gin.Context, db *gorm.DB) {
	// Retrieve the user ID from the JWT token
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Retrieve the task ID from the URL parameter
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	// Parse the task ID from the URL parameter
	var taskIDInt uint
	if _, err := fmt.Sscan(taskID, &taskIDInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Retrieve the task from the database
	var task models.Task
	if err := db.First(&task, taskIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Check if the user owns the task
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this task"})
		return
	}

	// Update the task with the new data
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task in the database
	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	// Return the updated task
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}

// DeleteTask handler for deleting a task
func DeleteTask(c *gin.Context, db *gorm.DB) {
	// Retrieve the user ID from the JWT token
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Retrieve the task ID from the URL parameter
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	// Parse the task ID from the URL parameter
	var taskIDInt uint
	if _, err := fmt.Sscan(taskID, &taskIDInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Retrieve the task from the database
	var task models.Task
	if err := db.First(&task, taskIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Check if the user owns the task
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this task"})
		return
	}

	// Delete the task from the database
	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// GetTaskByID handler for retrieving a task by ID
func GetTaskByID(c *gin.Context, db *gorm.DB) {
	// Retrieve the user ID from the JWT token
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Retrieve the task ID from the URL parameter
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	// Parse the task ID from the URL parameter
	var taskIDInt uint
	if _, err := fmt.Sscan(taskID, &taskIDInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Retrieve the task from the database
	var task models.Task
	if err := db.First(&task, taskIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Check if the user owns the task
	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to view this task"})
		return
	}

	// Return the task
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// Helper function to retrieve the user ID from the JWT token
func GetUserIDFromToken(c *gin.Context) (uint, error) {
	userToken := c.GetHeader("Authorization") // Get the authorization header
	if userToken == "" {
		return 0, fmt.Errorf("no authorization header provided")
	}

	// Split the header to remove the "Bearer " prefix
	tokenString := strings.Split(userToken, " ")[1]

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retrieve the JWT secret from the environment
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID in token")
	}
	return uint(userID), nil // Return the user ID as a uint
}
