package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"post/internal/database/models"
)

var Active models.User
var err error

var ConnectionString = "user=postgres password=akeceqm dbname=society sslmode=disable"

func InitDb(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return nil, err
	}
	return db, nil
}
