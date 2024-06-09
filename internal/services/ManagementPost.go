package services

import (
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
	"strconv"
	"time"
)

func CreatePost(title, text string, authorId int, db *sqlx.DB) (models.Post, error) {
	post := models.Post{}

	_, err := db.Exec("SELECT * FROM public.user WHERE id = $1", strconv.Itoa(authorId))
	if err != nil {
		return models.Post{}, err
	}

	post.Id = GenerateId()
	post.Title = title
	post.Text = text
	post.DateCreated = time.Now()
	post.AuthorId = authorId

	_, err = db.Exec("INSERT INTO post (id, title, text, date_created, author_id) VALUES ($1, $2, $3, $4, $5)", post.Id, post.Title, post.Text, post.DateCreated, post.AuthorId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetPost(db *sqlx.DB) ([]models.Post, error) {
	post := []models.Post{}

	err := db.Select(&post, "SELECT * FROM post")
	if err != nil {
		return post, err
	}
	return post, nil
}
