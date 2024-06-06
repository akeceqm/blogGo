package models

import "time"

type User struct {
	Name     string
	Email    string
	Password string
	Ip       string
	Data     time.Time
}

type IPResponse struct {
	IP string
}
