package services

import (
	"fmt"
	"post/internal/middlewares"
)

func MainFunction() {
	choise := ""
	fmt.Printf("\nОсновные действия в приложении\n [1] Просмотреть профиль\n [2] Добавить пост\n [3] Просмотреть пост(ы)\n [4]Выйти из аккаунта\n")
	fmt.Print("Введите номер действия чтобы продолжить: ")
	fmt.Scan(&choise)

	switch choise {
	case "1":
		middlewares.ClearScreen()
		MyProfile()
	case "2":
		// Здесь ваша логика для добавления поста
	case "3":
		// Здесь ваша логика для просмотра постов
	case "4":
		middlewares.ClearScreen()
		MainApp("Вы вышли из аккаунта")
	}
}
