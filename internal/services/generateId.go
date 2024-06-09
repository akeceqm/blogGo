package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GenerateId() (string, error) {
	password, err := middlewares.GenerateString(idLenght, idSet)
	if err != nil {
		fmt.Println("Не удалось сгенерировать id:", err)
		return "", err
	}
	return password, nil
}
