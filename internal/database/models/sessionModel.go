package models

import "time"

type Session struct {
	SessionID  string    `json:"session_id" db:"session_id"`
	UserID     string    `json:"user_id" db:"user_id"`
	Created_At time.Time `json:"created_at" db:"created_at"`
}
