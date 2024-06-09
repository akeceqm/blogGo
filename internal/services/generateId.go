package services

import (
	"fmt"
	"post/internal/middlewares"
)

func GenerateId(id []byte) error {
	if err = middlewares.GenerateString(idLenght, id, idSet); err != nil {
		fmt.Println("Не удалось сгенерировать id")
		return err
	}
	return nil
}
