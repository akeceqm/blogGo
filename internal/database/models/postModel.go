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

type FullPost struct {
	Id                string
	Title             string
	Text              string
	DateCreated       time.Time `db:"date_created" json:"date_created"`
	DateCreatedFormat string
	AuthorName        string `db:"nick_name" json:"nick_name"`
	AuthorId          string `db:"author_id" json:"author_id"`
	Comments          []FullComment
	CommentsCount     int
}

type FullComment struct {
	Id                string
	Text              string
	DateCreated       time.Time `db:"date_created" json:"date_created"`
	DateCreatedFormat string
	AuthorId          string `db:"author_id" json:"author_id"`
	PostId            string `db:"post_id" json:"post_id"`
	AuthorName        string `db:"nick_name" json:"nick_name"`
}
