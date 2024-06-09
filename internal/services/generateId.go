package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GenerateId() string {
	if id, err = middlewares.GenerateString(idLenght, idSet); err != nil {
		fmt.Println("Не удалось сгенерировать id " + err.Error())
		return err.Error()
	}
	return id
}
