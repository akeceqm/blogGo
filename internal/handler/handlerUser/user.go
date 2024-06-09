package handlerUser

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/middlewares"
	"post/internal/services"
	"time"
)

var err error

func GetUser(c *gin.Context, db *sqlx.DB) {
	var users []models.User

	// Выполняем запрос к базе данных для получения списка пользователей
	err := db.Select(&users, `SELECT * FROM public.user`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Отправляем список пользователей в формате JSON
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context, db *sqlx.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := services.GeneratePassword()
	currentTime := time.Now()
	_, err := db.Exec(`INSERT INTO public.user (id, login,email, password_hash,id_addres,date_registration) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, services.GenerateId(), user.Login, user.Email, middlewares.PasswordHash(password), middlewares.GetApi(), currentTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
