package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"post/cmd"
	"post/internal/database"
	"post/internal/handler"
)

func main() {
	cmd.Server = gin.Default()
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return
	}
	defer db.Close()

	handler.InitRoutes(cmd.Server, db)

	err = StartMain(cmd.Server)
	if err != nil {
		log.Fatalln("Неудачный запуск сервера")
	}
}

func StartMain(server *gin.Engine) error {
	log.Println("Сервер запущен")
	return server.Run(":8080")
}
