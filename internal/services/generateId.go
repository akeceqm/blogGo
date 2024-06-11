package services

import (
	"post/internal/middlewares"
)

func GenerateId() string {
	id := middlewares.GenerateString(idLenght, idSet)
	return id
}
