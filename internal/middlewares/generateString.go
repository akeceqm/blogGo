package middlewares

import (
	"fmt"
	"math/rand"
)

func GenerateString(length int, charset string) string {
	result := make([]byte, length)
	if charset == "" {
		fmt.Println("charset пустой  " + err.Error())
		return err.Error()
	}
	if length <= 0 {
		fmt.Println("length пустая " + err.Error())
		return err.Error()
	}
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
