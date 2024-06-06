package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"
)

func CreatePost() error {

	var Post models.Post
	fmt.Print("Введите название поста: ")
	title := middlewares.ReadString()
	if err != nil {
		fmt.Println(err)
		return err
	}

	Post.Title = title

	fmt.Print("Введите описание поста: ")
	description := middlewares.ReadString()
	if err != nil {
		fmt.Println(err)
		return err
	}

	Post.Description = description
	Post.Data = time.Now()

	Post.Author = database.GetActive()

	database.Posts = append(database.Posts, Post)

	database.PosMytView()

	MainFunction()
	return nil
}
