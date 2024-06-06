package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

type UserHandler interface {
	AuthorizationUser() error
	RegistrationUser() error
}

func MainApp(messageToUser string) {
	middlewares.ClearScreen()
	fmt.Println(messageToUser)

	for {
		choise := ""
		fmt.Println("[1] Регистрация\n[2] Авторизация\n[3] Просмотреть посты пользователей\n[4] Закрыть приложение")
		fmt.Print("Введите номер действия чтобы продолжить: ")
		fmt.Scan(&choise)

		switch choise {
		case "1":
			RegistrationUser()

		case "2":
			AuthorizationUser()
		case "3":
			database.PostAllView()
		case "4":
			fmt.Println("Закрытие приложения...")
			return
		default:
			middlewares.ClearScreen()
			fmt.Println("Введите правильный ответ")
		}
	}
}
