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

func GetUsers(c *gin.Context, db *sqlx.DB) {
	var users []models.User
	err := db.Select(&users, `SELECT *FROM public.user`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context, db *sqlx.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := services.GeneratePassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate password"})
		return
	}

	id, err := services.GenerateId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate id"})
		return
	}

	currentTime := time.Now()
	_, err = db.Exec(`INSERT INTO public.user (id, login,email, password_hash,id_address,date_registration) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, id, user.Email, user.Login, middlewares.PasswordHash(password), middlewares.GetApi(), currentTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
