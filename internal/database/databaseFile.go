package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"post/internal/database/models"
)

var Active models.User
var err error

var ConnectionString = "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=society sslmode=disable"

func InitDb(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд, " + err.Error())
		return nil, err
	}
	return db, nil
}
