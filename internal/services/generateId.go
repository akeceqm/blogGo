package services

import "post/internal/middlewares"

func GenerateId() string {
	return middlewares.GenerateString(idLenght, idSet)
}
