package mainApp

import (
	"fmt"
	"post/internal/database"
	"post/internal/services"
)

var (
	name           string
	email          string
	password       string
	repeatPassword string
	passwordLength int = 8
	err            error
)

func ChangePassword() {
	fmt.Print("Введите новый пароль: ")
	fmt.Scan(&password)

	passwordChecker := &services.CheckPasswordUser{}
	err = passwordChecker.CheckPassword(password, repeatPassword, passwordLength)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Повторите новый пароль: ")
	fmt.Scan(&repeatPassword)

	validateChecker := &services.CheckValidatePassword{}
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
