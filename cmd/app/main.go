package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"post/internal/database"
)

var server *gin.Engine
var db *sqlx.DB
var err error

func main() {
	server = gin.Default()
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return
	}
	defer db.Close()

	StartMain(server)
}

func StartMain(server *gin.Engine) error {
	if err := server.Run(":8080"); err != nil {
		log.Fatalln("Сервер не запущен!")
		return err
	}
	log.Println("Сервер успещно запущен!")
	return nil
}
