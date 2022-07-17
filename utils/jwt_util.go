package utils

import (
	"go_code/configs"
	"go_code/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

func GetToken(user models.User) (string, error) {
	claims := &CustomClaims{user.Username, user.Email, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Id:        user.Id.String(),
	}}

	// Create token with claims
	tokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	token, err := tokenGen.SignedString([]byte(configs.GetJWTKey()))
	if err != nil {
		log.Println("token generate error ", err.Error())
		return "", err
	}
	return token, nil
}
