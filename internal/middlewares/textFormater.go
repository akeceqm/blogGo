package middlewares

import "strings"

func TextFormatter(post struct{ Text string }) string {
	var formattedText string
	var currentLine string
	var count int

	words := strings.Fields(post.Text)

	for _, word := range words {
		if count+len(word) > 50 {
			formattedText += currentLine + "\n"
			currentLine = word + " "
			count = len(word) + 1
		} else {
			currentLine += word + " "
			count += len(word) + 1
		}
	}

	// Добавляем оставшуюся часть текста
	if len(currentLine) > 0 {
		formattedText += currentLine
	}

	// Проверяем последний символ и добавляем \n если нужно
	if formattedText[len(formattedText)-1] != ' ' {
		formattedText += "\n"
	}

	return formattedText
}
