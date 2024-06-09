package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GeneratePassword() string {

	if password, err = middlewares.GenerateString(passwordLenght, passwordSet); err != nil {
		fmt.Println("Не удалось сгенерировать id " + err.Error())
		return err.Error()
	}
	return password
}
