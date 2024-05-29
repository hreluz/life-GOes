package model

import "github.com/golang-jwt/jwt/v5"

// Login
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim (body token)
type Claim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
