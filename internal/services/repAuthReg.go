package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/database/models"
	"post/internal/middlewares"

	"golang.org/x/crypto/bcrypt"
)

func AuthorizationUser() error {
	database.GetAllUsersCopy()

	fmt.Print("Введите свой email: ")
	fmt.Scan(&email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var user *models.User
	if len(database.Users) == 0 {
		middlewares.ClearScreen()
		fmt.Println("в нашей соц-сети отсутсвуют юзеры.\nБудьте первыми!\n")
		return err
	} else {
		for i := range database.Users {
			if database.Users[i].Email == email {
				user = &database.Users[i]
			} else {
				fmt.Println("Такого email нет в базе данных!")
				return err
			}
		}
	}
	fmt.Print("Введите свой пароль: ")
	fmt.Scan(&password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Неверный пароль")
		return err
	}
	middlewares.ClearScreen()
	fmt.Printf("Пользователь %s успешно авторизован!\n", user.Name)
	MainFunction()
	return nil
}

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
	// Добавление пользователя
	database.Users = append(database.Users, User)
	middlewares.ClearScreen()
	fmt.Printf("name: %s \nemail: %s \npassword: %s \nip: %s \n", User.Name, User.Email, User.Password, User.Ip)
	return nil
}
