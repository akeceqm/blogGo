package middlewares

import (
	"fmt"
	"math/rand"
)

func GenerateString(length int, sqlParameter []byte, charset string) error {
	sqlParameter = make([]byte, length)

	if charset == "" {
		fmt.Println("charset пустой")
		return err
	}

	if length <= 0 {
		fmt.Println("length пустая")
		return err
	}

	for i := range sqlParameter {
		sqlParameter[i] = charset[rand.Intn(len(charset))]
	}

	return nil
}
