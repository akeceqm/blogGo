package middlewares

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) ([]byte, error) {

	if password == "" {
		fmt.Println("password пустая")
		return []byte{}, err
	}

	passwordHash := []byte(password)
	cost := 12
	hash, _ := bcrypt.GenerateFromPassword(passwordHash, cost)
	return hash, nil
}
