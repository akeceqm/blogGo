package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
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
func AuthMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "необходимо авторизоваться"})
			c.Abort()
			return
		}

		session, err := GetSessionByID(db, cookie)
		if err != nil || session.UserID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "необходимо авторизоваться"})
			c.Abort()
			return
		}

		c.Set("userID", session.UserID)
		c.Next()
	}
}
