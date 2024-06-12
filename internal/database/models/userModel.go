package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id               string         `db:"id"`
	Login            string         `db:"login" json:"login"`
	Email            string         `db:"email"`
	PasswordHash     string         `db:"password_hash" json:"password_hash"`
	Ip               sql.NullString `db:"ip_address"`
	DateRegistration time.Time      `db:"date_registration"`
	NickName         string         `db:"nick_name"`
	Description      string
}

type IPResponse struct {
	IP string
}
