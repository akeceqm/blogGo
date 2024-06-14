package middlewares

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateSessionID() (string, error) {
	// Генерируем 32 случайных байта
	b := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}

	// Кодируем в base64 и возвращаем как строку
	return base64.URLEncoding.EncodeToString(b), nil
}
