package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id               string         `db:"id" json:"id"`
	Name             string         `db:"nick_name" json:"nick_name"`
	Login            string         `db:"login" json:"login"`
	Email            string         `db:"email" json:"email"`
	PasswordHash     string         `db:"password_hash" json:"password_hash"`
	Ip               sql.NullString `db:"ip_address" json:"ip_address"`
	DateRegistration time.Time      `db:"date_registration" json:"dateRegistration"`
	Description      sql.NullString `db:"description" json:"description"`

}

type IPResponse struct {
	IP string
}
