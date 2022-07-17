package models

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}