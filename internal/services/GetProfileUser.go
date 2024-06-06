package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

func MyProfile() {
	choise := ""
	User, err := database.GetByUserEmail(email)
	if err != nil {
		fmt.Println("ошибка при получении данных пользователя:", err)
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

func ChangePassword() {
	fmt.Print("Введите новый пароль: ")
	fmt.Scan(&password)

	passwordChecker := &CheckPasswordUser{}
	err = passwordChecker.CheckPassword(password, passwordLength)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Повторите новый пароль: ")
	fmt.Scan(&repeatPassword)

	validateChecker := &CheckValidatePassword{}
	err = validateChecker.CheckValidate(password, repeatPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := database.GetByUserEmail(email) // Получаем пользователя по электронной почте
	if err != nil {
		fmt.Println("Ошибка при получении пользователя:", err)
		return
	}

	// Обновляем пароль пользователя
	err = database.UpdateUserPassword(user, password)
	if err != nil {
		fmt.Println("Ошибка при обновлении пароля:", err)
		return
	}

	fmt.Println("Пароль успешно изменен.")
	MyProfile()
}
