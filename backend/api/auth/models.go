package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Email string `json:"email"`
}
