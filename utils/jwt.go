package utils

import (
	"time"

	"api-gateway/config"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Role  string   `json:"role"`
	Scope []string `json:"scope"`
	jwt.RegisteredClaims
}

func GenerateGuestToken() (string, error) {
	claims := CustomClaims{
		Role:  "guest",
		Scope: []string{"content:read"},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.JWTIssuer,
			Audience:  []string{config.JWTAudience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "guest",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}
