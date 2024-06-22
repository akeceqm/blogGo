package handlerUser

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
)

func GetHandleUsers(c *gin.Context, db *sqlx.DB) {
	user, err := services.GetUser(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func PUTHandleUser(c *gin.Context, db *sqlx.DB) (models.User, error) {
	userId := c.Param("userId")
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return models.User{}, err
	}

	nick := newUser.NickName
	if nick.String == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "nickname не может быть пустым"})
		return models.User{}, errors.New("nickname не может быть пустым")
	}
	description := newUser.Description
	if description.String == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "decription не может быть пустым"})

		return models.User{}, errors.New("description не может быть пустым")
	}
	avatar := newUser.Avatar

	user, err := services.UpdateUser(db, userId, nick.String, description.String, avatar.String, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return models.User{}, err
	}
	c.JSON(http.StatusOK, user)

	return user, nil
}
