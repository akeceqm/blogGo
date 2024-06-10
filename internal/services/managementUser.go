package services

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"post/internal/database/models"
	"post/internal/middlewares"
	"time"
)

func GetUser(db *sqlx.DB) ([]models.User, error) {
	users := []models.User{}

	err := db.Select(&users, "SELECT * FROM public.user")
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUserByLogin(db *sqlx.DB, login string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, `SELECT * FROM public.user WHERE login = $1 `, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ошибка login или password")
		} else {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByCheckPassword(hashedPwd, plainPwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return errors.New("пароли не совпадают")
	}
	return nil
}

func PostUser(db *sqlx.DB, email string) (models.User, error) {
	var user models.User

	user.Id = GenerateId()
	user.Login = GenerateLogin()
	user.Email = email
	user.PasswordHash = GeneratePassword()
	user.DateRegistration = time.Now()

	_, err := db.Exec(`INSERT INTO public.user (id, login,email, password_hash,ip_address, date_registration) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, user.Id, user.Login, user.Email, middlewares.PasswordHash(user.PasswordHash), middlewares.GetApi(), user.DateRegistration)
	if err != nil {
		return user, errors.New("Неудачная регистрация. Попробуйте еще раз!")
	}

	return user, nil
}
