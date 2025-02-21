package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5" // Correct import
    "net/http"
    "strings"
)

// AuthMiddleware validates JWT tokens and protects routes
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        // Extract the token from the header (format: "Bearer <token>")
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        // Parse the token
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        // Handle token parsing errors
        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            }
            c.Abort()
            return
        }

        // Check if the token is valid
        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
            c.Abort()
            return
        }

        // Set the username in the context for use in protected routes
        c.Set("username", claims.Username)
        c.Next()
    }
}