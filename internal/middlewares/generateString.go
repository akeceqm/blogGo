package middlewares

import (
	"fmt"
	"math/rand"
)

func GenerateString(length int, charset string) string {

	var result []byte
	if charset == "" {
		fmt.Println("charset пустой  " + err.Error())
		return err

	}

	if length <= 0 {
		fmt.Println("length пустая " + err.Error())
		return "", err
	}

	for i := 0; i < length; i++ {
		result = append(result, charset[rand.Intn(len(charset))])
	}

	return string(result), nil
}
