package services

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(db *sqlx.DB) ([]models.User, error) {
	users := []models.User{}

	err := db.Select(&users, `SELECT * FROM public.user`)
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUserById(db *sqlx.DB, userId string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, nick_name, date_registration, description FROM public.user WHERE id = $1`
	err := db.Get(user, query, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByLogin(db *sqlx.DB, login string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, `SELECT * FROM public.user WHERE login = $1`, login)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Пользователь с логином %s не найден", login)
			return nil, errors.New("ошибка: неверный логин или пароль")
		} else {
			log.Printf("Ошибка при получении пользователя по логину: %v", err)
			return nil, err
		}
	}
	return &user, nil
}

func GetUserByCheckPassword(hashedPwd, plainPwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return errors.New("пароли не совпадают")
	}
	return nil
}

func PostUser(db *sqlx.DB, email string, name string) (models.User, error) {
	var user models.User

	user.Id = GenerateId()
	user.Login = GenerateLogin()
	user.Name = name
	user.Email = email
	user.PasswordHash = GeneratePassword()
	user.DateRegistration = time.Now()
	_, err := db.Exec(`INSERT INTO public.user (id,nick_name, login,email, password_hash,ip_address, date_registration) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`, user.Id, user.Name, user.Login, user.Email, middlewares.PasswordHash(user.PasswordHash), middlewares.GetApi(), user.DateRegistration)
	if err != nil {
		return user, errors.New("Неудачная регистрация. Попробуйте еще раз!")
	}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	return user, nil
}

func GetUserImage(c *gin.Context, db *sqlx.DB) {
	userId := c.Param("userId")

	var imagePath string
	query := "SELECT avatar FROM public.user WHERE id = $1"
	err := db.Get(&imagePath, query, userId)
	if err != nil || imagePath == "" {
		// Путь к дефолтному изображению
		defaultImagePath := "post/src/img/avatar.svg"
		c.File(defaultImagePath)
		return
	}

	c.File(imagePath)
}
