package services

import "post/internal/middlewares"

func GeneratePassword() string {
	return middlewares.GenerateString(passwordLenght, passwordSet)
}
