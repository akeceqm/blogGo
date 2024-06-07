package services

import (
	"fmt"
	"post/internal/database"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"
)

func CreateComment(id string) error {
	var Comment models.Comment

	Comment.PostId = id

	Comment.Id = middlewares.GenerateId(6)
	fmt.Print("Введите коментарий ")
	description = middlewares.ReadString()

	Comment.Description = description

	Comment.Data = time.Now()
	Comment.Author = database.GetActive()
	database.Comments = append(database.Comments, Comment)

	for index, val := range database.Posts {
		if val.Id == Comment.PostId {
			database.Posts[index].Comment = append(database.Posts[index].Comment, Comment)
		}
	}

	MainFunction()
	return nil
}
