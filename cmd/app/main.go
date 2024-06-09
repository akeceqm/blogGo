package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"post/cmd"
	"post/internal/database"
)

func main() {
	cmd.Server = gin.Default()
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return
	}
	defer db.Close()

	StartMain(cmd.Server)
}

func StartMain(server *gin.Engine) error {
	if err := server.Run(":8080"); err != nil {
		log.Fatalln("Сервер не запущен!")
		return err
	}
	log.Println("Сервер успещно запущен!")
	return nil
}
