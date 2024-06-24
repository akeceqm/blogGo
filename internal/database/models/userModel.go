package models

import "time"

type User struct {
	Id               string    `json:"id" db:"id"`
	NickName         string    `json:"nick_name" db:"nick_name"`
	Login            string    `json:"login" db:"login"`
	Email            string    `json:"email" db:"email"`
	PasswordHash     string    `json:"password_hash" db:"password_hash"`
	Avatar           string    `json:"avatar" db:"avatar"`
	DateRegistration time.Time `json:"date_registration" db:"date_registration"`
	Description      string    `json:"description" db:"description"`
	IPAddress        string    `json:"ip_address" db:"ip_address"`
}

type IPResponse struct {
	IP string
}
