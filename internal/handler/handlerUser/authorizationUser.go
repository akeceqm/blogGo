package handlerUser

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/middlewares"
	"post/internal/services"
)

func PostHandleAuthorizationUser(c *gin.Context, db *sqlx.DB) {
	var loginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password_hash"`
	}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	dbUser, err := services.GetUserByLogin(db, loginRequest.Login)
	if err != nil {
		if err.Error() == "ошибка: неверный логин или пароль" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "логин не найден"})
		} else {
			c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		}
		return
	}

	if err := services.GetUserByCheckPassword(dbUser.PasswordHash, loginRequest.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный пароль"})
		return
	}

	sessionID, err := middlewares.GenerateSessionID()
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	if err := services.CreateSession(db, sessionID, dbUser.Id); err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	cookie := http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{"id": dbUser.Id})
}

func GetHandleUserById(c *gin.Context, db *sqlx.DB) {
	userId := c.Param("userId")

	user, err := services.GetUserById(db, userId)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	registrationDate := user.DateRegistration.Format("2006-01-02")
	avatar := ""
	if user.Avatar.Valid {
		avatar = user.Avatar.String
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                user.Id,
		"nick_name":         user.NickName,
		"date_registration": registrationDate,
		"description":       user.Description.String,
		"avatar":            avatar,
	})
}
