package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var ConnectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=society sslmode=disable"

func InitDb(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Неудаченая попытка соединения с бд")
		return nil, err
	}
	return db, nil
}
