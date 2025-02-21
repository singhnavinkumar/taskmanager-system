package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5" // Correct import
    "time"
)

// Signup handles user registration
func Signup(c *gin.Context) {
    // For now, just return a success message
    c.JSON(200, gin.H{"message": "User signed up successfully"})
}

// Login handles user login and JWT token generation
func Login(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password") // Password is now used for validation

    // For simplicity, assume credentials are valid
    // In a real application, validate the username and password against a database
    if username == "" || password == "" {
        c.JSON(400, gin.H{"error": "Username and password are required"})
        return
    }

    // Create JWT token
    expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return the token to the client
    c.JSON(200, gin.H{"token": tokenString})
}