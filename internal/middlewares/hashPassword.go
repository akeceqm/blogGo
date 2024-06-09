package middlewares

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {

	if password == "" {
		fmt.Println("password пустая " + err.Error())
		return err.Error()
	}

	passwordHash := []byte(password)
	cost := 12
	hash, _ := bcrypt.GenerateFromPassword(passwordHash, cost)
	return string(hash)
}
