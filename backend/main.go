package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// User struct
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}

func main() {
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
        log.Fatal("Database connection failed:", err)
    }

    // Automatically migrate the User schema
    if err := db.AutoMigrate(&User{}); err != nil {
        log.Fatal("Failed to migrate database schema:", err)
    }

    // Setup Gin router
    router := gin.Default()
    router.POST("/register", func(c *gin.Context) {
        registerHandler(c, db)
    })
    router.POST("/login", func(c *gin.Context) {
        loginHandler(c, db)
    })

    // Run the Gin server on port 8080
    router.Run(":8080")
}

// Handler for user registration
func registerHandler(c *gin.Context, db *gorm.DB) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if username already exists BEFORE starting transaction
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// Username already exists
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = string(hashedPassword)

    // Begin transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback() // Rollback on error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit() // Commit transaction if successful
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Handler for user login
func loginHandler(c *gin.Context, db *gorm.DB) {
    var loginUser User
    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user from the database
    var dbUser User
    if err := db.Where("username = ?", loginUser.Username).First(&dbUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Compare passwords
    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  dbUser.ID,
        "username": dbUser.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    // Sign the token
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

