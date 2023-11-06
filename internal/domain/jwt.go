package domain

// import "github.com/golang-jwt/jwt"

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}
