package routes

import (
	"todo-list-go/backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures the Gin router with all API routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", func(c *gin.Context) { handlers.Register(c, db) })
		auth.POST("/login", func(c *gin.Context) { handlers.Login(c, db) })
	}

	// Task routes
	task := router.Group("/tasks")
	{
		task.POST("/", func(c *gin.Context) { handlers.CreateTask(c, db) })
		task.GET("/", func(c *gin.Context) { handlers.GetTasks(c, db) })
		task.GET("/:id", func(c *gin.Context) { handlers.GetTaskByID(c, db) })
		task.PUT("/:id", func(c *gin.Context) { handlers.UpdateTask(c, db) })
		task.DELETE("/:id", func(c *gin.Context) { handlers.DeleteTask(c, db) })
	}
}
