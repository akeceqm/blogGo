package models

import "time"

type User struct {
	Id               string    `db:"id"`
	Login            string    `db:"login"`
	Email            string    `db:"email"`
	PasswordHash     string    `db:"password_hash"`
	Ip               string    `db:"id_address"`
	DateRegistration time.Time `db:"date_registration"`
}

type IPResponse struct {
	IP string
}
