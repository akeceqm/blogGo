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
var err error

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
			fmt.Printf("Id: %s\nНазвание: %s\nОписание: %s\nДата публикации: %s\nКомментарев: %d\n", val.Id, val.Title, val.Description, val.Data.Format("02.01.2006 15:04"), len(val.Comment))
		}
	}
}

func GetLenMyPost() int {
	var count = 0
	for _, val := range Posts {
		if val.Author == GetActive() {
			for i := 0; i < len(val.Comment); i++ {
				count++
			}
		}
	}
	return count
}

func PostAllView() {
	totalPosts := len(Posts)
	if totalPosts > 0 {
		for i, val := range Posts {
			fmt.Printf("Название: %s\nОписание: %s\nДата публикации: %s\nКоментариев: %d\n *Автор: %s\n", val.Title, val.Description, val.Data.Format("02.01.2006 15:04"), len(val.Comment), val.Author.Name)
			if i < totalPosts-1 {
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Нет доступных постов.")
	}
}

func PutPost(id, title, description string) {
	for index, val := range Posts {
		if val.Id == id {
			Posts[index].Title = title
			if title == "" {
				fmt.Println("Ошибка! Введите хотя бы 1 символ")
				break
			}

			Posts[index].Description = description

			if description == "" {
				fmt.Println("Ошибка! Введите хотя бы 1 символ")
				break
			}

		}
	}
}

func DelPost(id string) error {
	for indx, val := range Posts {
		if val.Id == id {
			Posts = append(Posts[:indx], Posts[indx+1:]...)
			return nil
		}
	}
	return errors.New("Введи корректный id")
}

func GetPostsForComments() {
	for _, val := range Posts {
		fmt.Printf("Id: %s\nНазвание: %s\n", val.Id, val.Title)
		return
	}
}

func GetComment() {
	for _, val := range Comments {
		fmt.Printf("Id: %s\nId поста: %s\nТекст комментария: %s\nАвтор коментария: %s", val.Id, val.PostId, val.Description, val.Author.Name)
		return
	}
}
