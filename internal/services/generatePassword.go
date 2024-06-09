package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GeneratePassword(password []byte) error {
	if err = middlewares.GenerateString(passwordLenght, password, passwordSet); err != nil {
		fmt.Println("Не удалось сгенерировать id")
		return err
	}
	return nil
}
