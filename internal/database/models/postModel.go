package models

import "time"

type Post struct {
	Id          int
	Title       string
	Description string
	Data        time.Time
	Author      string
	IdUser      int
}

type Comment struct {
	Id            int
	Description   string
	Data          time.Time
	Author        string
	PostConnected string
	IdUser        int
}

type Lice struct{}
