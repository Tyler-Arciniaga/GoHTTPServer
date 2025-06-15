package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ContextKey string

const UserIDKey ContextKey = "userID"
