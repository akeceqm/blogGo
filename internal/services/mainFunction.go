package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

func MainFunction() {
	for {
		choiсe := ""
		fmt.Printf("\nОсновные действия в приложении\n [1] Просмотреть профиль\n [2] Добавить пост\n [3] Просмотреть мой(и) пост(ы)\n [4] Просмотреть все посты\n [5] Выйти из аккаунта\n")
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
			PostChange()

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

func PostChange() {
	var countPost = database.GetLenMyPost()
	for {
		if countPost == 0 {
			fmt.Println("Постов нету")
			break
		} else {
			var input = ""
			fmt.Println("[*] Изменить пост || [!] Удалить пост || [Любой символ отличнй от этих] Ничего не делать")
			fmt.Print("Введите ответ: ")
			fmt.Scan(&input)
			if input == "*" {
				fmt.Print("Введите id: ")
				fmt.Scan(&id)
				if id == "" {
					fmt.Println("Введите id")
					break
				}
				fmt.Print("Введите Название: ")
				fmt.Scan(&title)
				fmt.Print("Введите Описание: ")
				fmt.Scan(&description)
				database.PutPost(id, title, description)
			} else if input == "!" {
				for {
					fmt.Print("Введите id: ")
					fmt.Scan(&id)
					if id == "" {
						fmt.Println("Введите id")
						continue
					}
					err := database.DelPost(id)
					if err != nil {
						fmt.Println(err.Error())
						fmt.Print("Введите id: ")
						continue
					}
					fmt.Println("Все успещно удалено")
					break
				}
			} else {
				break
			}
		}
	}
}
