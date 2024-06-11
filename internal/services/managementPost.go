package services

import (
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
	"time"
)

func CreatePost(title, text string, authorId string, db *sqlx.DB) (models.Post, error) {
	post := models.Post{}

	err := db.Get("SELECT * FROM public.post WHERE id = $1", authorId)

	if err != nil {
		return models.Post{}, err
	}

	post.Id = GenerateId()
	post.Title = title
	post.Text = text
	post.DateCreated = time.Now()
	post.AuthorId = authorId

	_, err = db.Exec("INSERT INTO public.post (id, title, text, date, author_id) VALUES ($1, $2, $3, $4, $5)", post.Id, post.Title, post.Text, post.DateCreated, post.AuthorId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetPost(db *sqlx.DB) ([]models.Post, error) {
	post := []models.Post{}

	err := db.Select(&post, "SELECT * FROM public.post")
	if err != nil {
		return post, err
	}
	return post, nil
}

func GetPostsByAuthorId(idAuthor string, db *sqlx.DB) ([]models.Post, error) {
	var post []models.Post

	err := db.Select(&post, "SELECT * FROM public.post WHERE author_id = $1", idAuthor)
	if err != nil {
		return []models.Post{}, err
	}
	return post, nil
}

func GetPostByBetweenDate(startDate, endDate time.Time, db *sqlx.DB) ([]models.Post, error) {
	var post []models.Post

	err := db.Select(&post, "SELECT * FROM public.post WHERE date_created BETWEEN $1 AND $2", startDate, endDate)
	if err != nil {
		return []models.Post{}, err
	}

	return post, nil
}

func UpdatePost(id string, title, text string, db *sqlx.DB) error {
	var postDB models.Post

	err := db.Get(&postDB, "SELECT * FROM public.post WHERE id = $1", id)
	if err != nil {
		return err
	}
	if title == "" {
		title = postDB.Title
	}
	if text == "" {
		text = postDB.Text
	}

	_, err = db.Exec("UPDATE public.post SET title = $1, text = $2 WHERE id = $3", title, text, id)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id string, db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM public.post WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
