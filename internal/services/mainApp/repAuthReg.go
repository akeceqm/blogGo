package mainApp

import (
	"errors"
	"fmt"
	"post/internal/database"
	"post/internal/database/models"
	"post/internal/middlewares"
	"post/internal/services"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerImpl struct{}

func (ah AuthHandlerImpl) AuthorizationUser() error {
	database.GetAllUsersCopy()

	fmt.Print("Введите свой email: ")
	fmt.Scan(&email)
	if err != nil {
		return err
	}

	var user *models.User
	for i := range database.Users {
		if database.Users[i].Email == email {
			user = &database.Users[i]
			break
		}
	}

	if user == nil {
		return errors.New("Такого email нет в базе данных!")
	}

	fmt.Print("Введите свой пароль: ")
	fmt.Scan(&password)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("Неверный пароль")
	}
	middlewares.ClearScreen()
	fmt.Printf("Пользователь %s успешно авторизован!\n", user.Name)
	MainFunction()
	return nil
}

type RegistrationUserItmpl struct{}

func (ru *RegistrationUserItmpl) RegistrationUser() error {
	database.GetAllUsersCopy()

	var User models.User
	fmt.Print("Введите свое имя: ")
	fmt.Scan(&name)
	User.Name = name

	fmt.Print("Введите свою почту: ")
	fmt.Scan(&email)

	emailChecker := &services.CheckEmailUser{}
	err = emailChecker.CheckEmail(email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	User.Email = email

	fmt.Print("Введите свой пароль: ")
	fmt.Scan(&password)

	passwordChecker := &services.CheckPasswordUser{}
	err = passwordChecker.CheckPassword(password, repeatPassword, passwordLength)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Print("Повторите свой пароль: ")
	fmt.Scan(&repeatPassword)

	validateChecker := &services.CheckValidatePassword{}
	err = validateChecker.CheckValidate(password, repeatPassword)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	User.Password = string(middlewares.PasswordHash(password))
	User.Ip = middlewares.GetApi()
	// Добавление пользователя
	database.Users = append(database.Users, User)
	middlewares.ClearScreen()
	fmt.Printf("name: %s \nemail: %s \npassword: %s \nip: %s \n", User.Name, User.Email, User.Password, User.Ip)
	return nil
}
