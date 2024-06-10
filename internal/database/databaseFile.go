package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)


var ConnectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=society sslmode=disable"


func InitDb(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд, " + err.Error())
		return nil, err
	}
	return db, nil
}
