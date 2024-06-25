package services

import (
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
	"time"
)

func CreateSession(db *sqlx.DB, sessionID, userID string) error {
	session := models.Session{
		SessionID:  sessionID,
		UserID:     userID,
		Created_At: time.Now(),
	}

	_, err := db.Exec(`INSERT INTO public.sessions (session_id, user_id, created_at) VALUES ($1, $2, $3)`, session.SessionID, session.UserID, session.Created_At)
	return err
}

func GetSessionByID(db *sqlx.DB, sessionID string) (*models.Session, error) {
	var session models.Session
	err := db.Get(&session, `SELECT * FROM public.sessions WHERE session_id = $1`, sessionID)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func DeleteSession(db *sqlx.DB, userId string) error {
	_, err := db.Exec(`DELETE FROM public.sessions WHERE user_id = $1`, userId)
	if err != nil {
		return err
	}
	return nil
}

//func AuthMiddleware(db *sqlx.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		cookie, err := c.Cookie("session_id")
//		if err != nil || cookie == "" {
//			c.Redirect(http.StatusFound, "/authorization") // Redirect to login if session cookie not found
//			c.Abort()
//			return
//		}
//
//		session, err := GetSessionByID(db, cookie)
//		if err != nil || session.UserID == "" {
//			c.Redirect(http.StatusFound, "/authorization") // Redirect to login if session invalid
//			c.Abort()
//			return
//		}
//
//		c.Set("userID", session.UserID)
//		c.Next()
//	}
//}

func IsUserAuthorized(db *sqlx.DB, userID string) (bool, error) {
	var user models.User
	err := db.Get(&user, "SELECT * FROM public.user WHERE id = $1", userID)
	if err != nil {
		return false, err
	}
	return true, nil // Предполагаем, что пользователь авторизован, если найден в базе данных
}
