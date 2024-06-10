package services

import (
	"post/internal/middlewares"
)

func GenerateLogin() string {
	login := middlewares.GenerateString(loginLenght, loginSet)
	return login
}
