package middlewares

import (
	"fmt"
	"math/rand"
)

func GenerateString(length int, charset string) (string, error) {
	if charset == "" {
		fmt.Println("Длина равна charset 0")
		return "", err
	}

	if length <= 0 {
		fmt.Println("Длина равна пароля 0")
		return "", err
	}

	result := make([]rune, length)
	for i := range result {
		result[i] = rune(charset[rand.Intn(len(charset))])
	}

	return string(result), nil
}
