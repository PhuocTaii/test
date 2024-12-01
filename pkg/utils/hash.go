package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPasswordStr := string(hashedPassword)
	return &hashedPasswordStr, nil
}

func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}