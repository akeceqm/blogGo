package services


import (
	"post/internal/middlewares"
)

func GenerateId() string {
	id := middlewares.GenerateString(passwordLenght, passwordSet)
	return id
}
