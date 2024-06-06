package database

import (
	"errors"
	"fmt"
	"post/internal/database/models"
	"post/internal/middlewares"
)

var Users []models.User
var Posts []models.Post
var Comments []models.Comment
var Active models.User

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

func GetActive() models.User {
	return Active
}

func DelActive() models.User {
	Active = models.User{}
	return Active
}

func PutActive(User models.User) {
	Active = User
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

func PosMytView() {
	for _, val := range Posts {
		if val.Author == GetActive() {
			fmt.Printf("Название: %s\nОписание: %s\nДата публикации: %s\n\n", val.Title, val.Description, val.Data.Format("02.01.2006 15:04"))
		}
	}
}

func PostAllView() {
	totalPosts := len(Posts)
	if totalPosts > 0 {
		for i, val := range Posts {
			fmt.Printf("Название: %s\nОписание: %s\nДата публикации: %s\n *Автор: %s\n", val.Title, val.Description, val.Data.Format("02.01.2006 15:04"), val.Author.Name)
			if i < totalPosts-1 {
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Нет доступных постов.")
	}
}
