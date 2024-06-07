package handler

import (
	"fmt"
	"post/internal/database"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"
)

func RegistrationUser() error {
	database.GetAllUsersCopy()

	var User models.User
	fmt.Print("Введите свое имя: ")
	fmt.Scan(&name)
	User.Name = name

	fmt.Print("Введите свою почту: ")
	fmt.Scan(&email)

	emailChecker := &CheckEmailUser{}
	err = emailChecker.CheckEmail(email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	User.Email = email

	fmt.Print("Введите свой пароль: ")
	fmt.Scan(&password)

	passwordChecker := &CheckPasswordUser{}
	err = passwordChecker.CheckPassword(password, passwordLength)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Print("Повторите свой пароль: ")
	fmt.Scan(&repeatPassword)

	validateChecker := &CheckValidatePassword{}
	err = validateChecker.CheckValidate(password, repeatPassword)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	User.Password = string(middlewares.PasswordHash(password))
	User.Ip = middlewares.GetApi()
	User.Data = time.Now()

	database.Users = append(database.Users, User)
	middlewares.ClearScreen()
	fmt.Printf("name: %s \nemail: %s \npassword: %s \nip: %s \nD", User.Name, User.Email, User.Password, User.Ip)
	return nil
}
