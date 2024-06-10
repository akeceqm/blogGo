package services

import (
	"post/internal/middlewares"
)

func GeneratePassword() string {
	password := middlewares.GenerateString(passwordLenght, passwordSet)
	return password
}
