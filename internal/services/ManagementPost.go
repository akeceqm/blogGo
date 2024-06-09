package services

import (
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"
)

func CreatePost(title, text string, authorId int, db *sqlx.DB) error {
	post := models.Post{}

	err := db.Get("SELECT * FROM post WHERE id = $1", string(authorId))
	if err != nil {
		return err
	}

	post.Id = middlewares.GenerateId()
	post.Title = title
	post.Text = text
	post.DateCreated = time.Now()
	post.AuthorId = authorId

	_, err = db.Exec("INSERT INTO post (id, title, text, date, author_id) VALUES ($1, $2, $3, $4, $5)", post.Id, post.Title, post.Text, post.DateCreated, post.AuthorId)
	if err != nil {
		return err
	}

	return nil
}

func GetPost(db *sqlx.DB) ([]models.Post, error) {
	post := []models.Post{}

	err := db.Select(&post, "SELECT * FROM post")
	if err != nil {
		return post, err
	}
	return post, nil
}
