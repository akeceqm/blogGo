package middlewares

import (
	"math/rand"
)

func GenerateString(length int, charset string) (string, error) {
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result), err
}
