package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 1. Bikin Struct Custom
type JWTCustomClaims struct {
	Role string `json:"role"` // Field tambahan kita
	jwt.RegisteredClaims      // Embbed field standar (Exp, Sub, dll)
}

func GenerateToken(userId string, role string, secret string) (string, error) {
	// 2. Isi data ke struct
	claims := JWTCustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId, // Setara "sub"
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Setara "exp"
			IssuedAt:  jwt.NewNumericDate(time.Now()), // Setara "iat"
		},
	}

	// 3. Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}