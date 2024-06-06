package services

import (
	"errors"
	"post/internal/database"
	"post/internal/middlewares"
	"strconv"
)

type CheckEmailUser struct{}

func (c *CheckEmailUser) CheckEmail(email string) error {
	for _, val := range database.Users {
		if val.Email == email {
			middlewares.ClearScreen()
			return errors.New("Такой email  уже есть: " + val.Email)
		}
	}
	return nil
}

type CheckPasswordUser struct{}

func (c *CheckPasswordUser) CheckPassword(password string, passwordLenght int) error {
	if len(password) < passwordLenght {
		str := strconv.Itoa(passwordLenght)
		middlewares.ClearScreen()
		return errors.New("пароль должен быть не меньше: " + str)
	}
	return nil
}

type CheckValidatePassword struct{}

func (c *CheckValidatePassword) CheckValidate(password, repeatPassword string) error {
	if repeatPassword != password {
		middlewares.ClearScreen()
		return errors.New("пароли не совпадают")
	}
	return nil
}
