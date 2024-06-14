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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в запросе"})
		return
	}

	dbUser, err := services.GetUserByLogin(db, loginRequest.Login)
	if err != nil {
		if err.Error() == "ошибка: неверный логин или пароль" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "логин не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "внутренняя ошибка сервера"})
		}
		return
	}

	if err := services.GetUserByCheckPassword(dbUser.PasswordHash, loginRequest.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	sessionID, err := middlewares.GenerateSessionID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "внутренняя ошибка сервера"})
		return
	}

	// Устанавливаем куку с идентификатором сессии
	cookie := http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	}
	http.SetCookie(c.Writer, &cookie)
}
func GetHandleUserById(c *gin.Context, db *sqlx.DB) {
	userId := c.Param("userId")

	// Вызываем сервис для получения данных пользователя из базы данных
	user, err := services.GetUserById(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "внутренняя ошибка сервера"})
		return
	}
	registrationDate := user.DateRegistration.Format("2006-01-02")
	// Возвращаем данные пользователя в ответе
	c.JSON(http.StatusOK, gin.H{
		"id":                user.Id,
		"nick_name":         user.Name,
		"date_registration": registrationDate,
		"description":       user.Description.String,
		"avatar":            user.Avatar,
	})
}
