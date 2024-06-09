package models

import (
	"time"
)

type Post struct {
	Id          string
	Title       string
	Text        string
	DateCreated time.Time `db:"date_created" json:"date_created"`
	AuthorId    int       `db:"author_id" json:"author_id"`
	Comment     []Comment
}

type Comment struct {
	Id         string
	Text       string
	DateCreate time.Time `db:"date_created" json:"date_created"`
	AuthorId   int       `db:"author_id" json:"author_id"`
	PostId     string
}
