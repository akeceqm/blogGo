package services

import (
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
	"sort"
	"time"
)

func GetCommentsByAuthorId(id string, db *sqlx.DB) ([]models.Comment, error) {
	var comment []models.Comment

	err := db.Select(&comment, "SELECT * FROM public.comment WHERE comment.author_id = $1", id)
	if err != nil {
		return []models.Comment{}, err
	}

	return comment, nil
}

func GetCommentsByPostId(id string, db *sqlx.DB) ([]models.Comment, error) {
	var comment []models.Comment

	err := db.Get(&models.Post{}, "SELECT * FROM public.post WHERE id = $1", id)
	if err != nil {
		return []models.Comment{}, err
	}

	err = db.Select(&comment, "SELECT * FROM public.comment WHERE comment.post_id = $1", id)
	if err != nil {
		return []models.Comment{}, err
	}

	sort.Slice(comment, func(i, j int) bool {
		return comment[i].DateCreated.After(comment[j].DateCreated)
	})

	return comment, nil
}

func GetCommentsByPostIdFull(id string, db *sqlx.DB) ([]models.FullComment, error) {
	var comment []models.FullComment

	err := db.Get(&models.Post{}, "SELECT * FROM public.post WHERE id = $1", id)
	if err != nil {
		return []models.FullComment{}, err
	}

	err = db.Select(&comment, "SELECT public.comment.*, public.user.nick_name FROM public.comment JOIN public.user ON comment.author_id = public.user.id WHERE comment.post_id = $1", id)
	if err != nil {
		return []models.FullComment{}, err
	}

	sort.Slice(comment, func(i, j int) bool {
		return comment[i].DateCreated.After(comment[j].DateCreated)
	})

	return comment, nil
}

func CreateComment(text string, authorId string, postId string, db *sqlx.DB) error {
	dateCreated := time.Now().Format(dateFormat)
	_, err := db.Exec("INSERT INTO public.comment (id, text, date_created, author_id, post_id) VALUES ($1, $2, $3, $4, $5)", GenerateId(), text, dateCreated, authorId, postId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(idComment string, db *sqlx.DB) error {
	var comment models.Comment
	err := db.Get(&comment, "SELECT * FROM public.comment WHERE id = $1", idComment)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM public.comment WHERE id = $1", idComment)
	if err != nil {
		return err
	}

	return nil
}

func UpdateComment(idComment string, text string, db *sqlx.DB) error {
	var comment models.Comment

	err := db.Get(&comment, "SELECT * FROM public.comment WHERE id = $1", idComment)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE public.comment SET text = $1 WHERE id = $2", text, idComment)
	if err != nil {
		return err
	}

	return nil
}
