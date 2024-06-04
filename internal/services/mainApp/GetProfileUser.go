package mainApp

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

func MyProfile() {
	choise := ""
	User, err := database.GetByUserEmail(email)
	if err != nil {
		fmt.Println("Ошибка при получении данных пользователя:", err)
		return
	}

	fmt.Printf("Name: %s \nEmail: %s\n[*] Сменить пароль\n[!] Выйти\n", User.Name, User.Email)
	fmt.Print("Введите вариант: ")
	fmt.Scan(&choise)

	switch choise {
	case "*":
		fmt.Println("Вы уверены что хотите сменить пароль?")
		fmt.Print("Если *да* то нажмите [1] если *нет* то нажмите [2]: ")
		fmt.Scan(&choise)
		switch choise {
		case "1":
			middlewares.ClearScreen()
			ChangePassword()
		case "2":
			MyProfile()
		default:
			middlewares.ClearScreen()
			fmt.Println("Введите правильный знак")
			MyProfile()
		}

	case "!":
		MainFunction()
	default:
		fmt.Println("Введите правильный знак")
	}
}
