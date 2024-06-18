package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id               string         `db:"id" json:"id"`
	NickName         string         `db:"nick_name" json:"nick_name"`
	Login            string         `db:"login" json:"login"`
	Email            string         `db:"email" json:"email"`
	PasswordHash     string         `db:"password_hash" json:"password_hash"`
	Ip               sql.NullString `db:"ip_address" json:"ip_address"`
	DateRegistration time.Time      `db:"date_registration" json:"date_registration"`
	Description      sql.NullString `db:"description" json:"description"`
	Avatar           sql.NullString `db:"avatar" json:"avatar"`
}

type IPResponse struct {
	IP string
}
