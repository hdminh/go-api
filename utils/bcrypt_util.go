package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncryptPassword(data string) ([]byte, error) {
	encoded, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err !=  nil {
		return nil, err
	}
	log.Println("Encode password ", data, " ", encoded)
	return encoded, nil
}

func CheckPassword(userPW, typePW []byte) bool {
	if err := bcrypt.CompareHashAndPassword(userPW, typePW); err != nil {
		return false
	}
	return true
}  