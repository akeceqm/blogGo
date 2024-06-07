package models

import (
	"time"
)

type Post struct {
	Id          string
	Title       string
	Description string
	Data        time.Time
	Author      User
	IdUser      int
	Comment     []Comment
}

type Comment struct {
	Id          int
	Description string
	Data        time.Time
	Author      string
	PostId      string
	UserId      int
}
