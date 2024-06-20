package services

import (
	"post/internal/middlewares"
)

func GenerateLogin() string {
	login := middlewares.GenerateString(loginLenght, loginSet)
	login = "sever-" + login
	return login
}
