package models

import "time"

type User struct {
	Name             string
	Email            string
	Password         string
	Ip               string
	DateRegistration time.Time `db:"date_registration" json:"date_registration"`
}

type IPResponse struct {
	IP string
}
