package middlewares

import (
	"math/rand"
	"time"
)

func GenerateId(length int) string {
	// символы, которые могут быть использованы в пароле
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	rand.Seed(time.Now().UnixNano())

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
