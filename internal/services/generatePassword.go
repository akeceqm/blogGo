package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GeneratePassword() (string, error) {
	password, err := middlewares.GenerateString(passwordLenght, passwordSet)
	if err != nil {
		fmt.Println("Не удалось сгенерировать пароль:", err)
		return "", err
	}
	return password, nil
}
