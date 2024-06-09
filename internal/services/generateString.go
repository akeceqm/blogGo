package services

import (
	"math/rand"
)

func GenerateString(length int, charset string) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
