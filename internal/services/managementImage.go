package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/internal/database/models"
)

func PostImageUser(db *sqlx.DB, userID string, image string, c *gin.Context) (*models.User, error) {
	_, err := db.Exec(`INSERT INTO public.user (id, avatar) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET avatar = $2`, userID, image)
	if err != nil {
		c.HTML(404, "400.html", err.Error())
	}
	user, err := GetUserById(db, userID)
	return user, nil
}

func GetImageUser(db *sqlx.DB, userId string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, `SELECT avatar FROM public.user WHERE id = $1`, userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func UpdateAvatar(db *sqlx.DB, userId, avatar string) (models.User, error) {
//
//}
