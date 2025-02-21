package auth

import "github.com/golang-jwt/jwt/v5" // Correct import

var jwtKey = []byte("your_jwt_secret") // Replace with a strong secret key

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}