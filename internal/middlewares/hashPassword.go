package middlewares

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) []byte {
	passwordHash := []byte(password)
	cost := 12
	hash, _ := bcrypt.GenerateFromPassword(passwordHash, cost)
	return hash
}
