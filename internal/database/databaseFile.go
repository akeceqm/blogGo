package database

import (
	"errors"
	"post/internal/database/models"
	"post/internal/middlewares"
)

var Users []models.User
var Post []models.Post
var Comment []models.Comment

func GetAllUsersCopy() *[]models.User {
	return &Users
}

func GetByUserEmail(email string) (*models.User, error) {
	for _, user := range Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}

func UpdateUserPassword(user *models.User, newPassword string) error {
	for i, u := range Users {
		if u.Email == user.Email {
			Users[i].Password = string(middlewares.PasswordHash(newPassword)) // Обновляем пароль у пользователя в списке
			return nil
		}
	}
	return errors.New("Пользователь не найден")
}
