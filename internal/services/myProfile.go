package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/middlewares"
)

func MyProfile() {
	for {
		choice := ""
		User, err := database.GetByUserEmail(database.GetActive().Email)
		if err != nil {
			fmt.Println("ошибка при получении данных пользователя:", err)
			return
		}

		fmt.Printf("Name: %s \nEmail: %s\nДата регистрации аккаунта: %s\n[*] Сменить пароль\n[!] Выйти\n", User.Name, User.Email, User.Data.Format("02.01.2006 15:04"))
		fmt.Print("Введите вариант: ")
		fmt.Scan(&choice)

		if choice == "*" {
			fmt.Println("Вы уверены что хотите сменить пароль?")
			fmt.Print("Если *да* то нажмите [1] если *нет* то нажмите [2]: ")
			fmt.Scan(&choice)

			if choice == "1" {
				middlewares.ClearScreen()
				ChangePassword()
			} else if choice == "2" {
				break
			} else {
				middlewares.ClearScreen()
				fmt.Println("Введите правильный знак")
			}
		} else if choice == "!" {
			MainFunction()
		} else {
			fmt.Println("Введите правильный знак")
		}
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

	user, err := database.GetByUserEmail(email)
	if err != nil {
		fmt.Println("Ошибка при получении пользователя:", err)
		return
	}

	err = database.UpdateUserPassword(user, password)
	if err != nil {
		fmt.Println("Ошибка при обновлении пароля:", err)
		return
	}

	fmt.Println("Пароль успешно изменен.")
	MyProfile()
}
