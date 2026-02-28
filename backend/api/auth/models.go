package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Email string `json:"email"`
}

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type registerPayload struct {
	Email     string `json:"email"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}
