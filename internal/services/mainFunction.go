package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

func MainFunction() {
	for {
		choiсe := ""
		fmt.Printf("\nОсновные действия в приложении\n [1] Просмотреть профиль\n [2] Добавить пост\n [3] Просмотретьм мой(и) пост(ы)\n [4] Просмотреть все посты\n [5] Выйти из аккаунта\n")
		fmt.Print("Введите номер действия чтобы продолжить: ")
		fmt.Scan(&choiсe)

		if choiсe == "1" {
			middlewares.ClearScreen()
			MyProfile()
		} else if choiсe == "2" {
			middlewares.ClearScreen()
			CreatePost()
		} else if choiсe == "3" {
			middlewares.ClearScreen()
			database.PosMytView()
		} else if choiсe == "4" {
			middlewares.ClearScreen()
			database.PostAllView()
		} else if choiсe == "5" {
			middlewares.ClearScreen()
			for _, val := range database.Users {
				if val.Email == email {
					database.DelActive()
				}
			}
			MainApp("Вы вышли из аккаунта")
		} else {
			fmt.Println("Неправильный выбор, попробуйте снова.")
		}

	}
}
