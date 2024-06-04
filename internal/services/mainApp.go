package services

import (
	"fmt"
	"post/internal/middlewares"
)

type UserHandler interface {
	AuthorizationUser() error
	RegistrationUser() error
}

func MainApp(messageToUser string) {
	middlewares.ClearScreen()
	fmt.Println(messageToUser)

	authHandler := &AuthHandlerImpl{}
	regHandler := &RegistrationUserItmpl{}
	for {
		choise := ""
		fmt.Println("[1] Регистрация\n[2] Авторизация\n[3]Просмотреть посты пользователей\n[4] Закрыть приложение")
		fmt.Print("Введите номер действия чтобы продолжить: ")
		fmt.Scan(&choise)

		switch choise {
		case "1":
			if err := regHandler.RegistrationUser(); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "2":
			if err := authHandler.AuthorizationUser(); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Авторизация прошла успешно!")
			}
		case "3":
			fmt.Println("Функция пока не реализована")
		case "4":
			fmt.Println("Закрытие приложения...")
			return
		default:
			middlewares.ClearScreen()
			fmt.Println("Введите правильный ответ")
		}
	}
}
