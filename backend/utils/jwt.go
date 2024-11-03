package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT token with the given user ID and username
func GenerateJWT(userID uint, username string) (string, error) {
	// Create JWT token
	expirationTime := time.Now().Add(time.Hour * 2).Unix() // Token expires in 2 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      expirationTime, // Token expires in 2 hours
	})

	// Sign the token using the JWT_SECRET from the environment
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	// Log the token and expiration time for debugging
	fmt.Printf("Generated Token: %s, Expires at: %d\n", tokenString, expirationTime)

	return tokenString, nil
}
