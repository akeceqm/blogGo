package handlerUser

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
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
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Логин не найден"})
		return
	}

	if dbUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Пользователь не найден"})
		return
	}
	if err := services.GetUserByCheckPassword(dbUser.PasswordHash, loginRequest.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dbUser)
}
