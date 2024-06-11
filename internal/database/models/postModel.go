package models

import (
	"time"
)

type Post struct {
	Id          string
	Title       string
	Text        string
	DateCreated time.Time `db:"date_created" json:"date_created"`
	AuthorId    string    `db:"author_id" json:"author_id"`
	Comment     []Comment
}

type Comment struct {
	Id          string
	Text        string
	DateCreated time.Time `db:"date_created" json:"date_created"`
	AuthorId    string    `db:"author_id" json:"author_id"`
	PostId      string    `db:"post_id" json:"post_id"`
}
