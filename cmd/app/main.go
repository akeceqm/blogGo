package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"post/cmd"
	"post/internal/database"
	"post/internal/handler/handlerUser"
)

func main() {
	cmd.Server = gin.Default()
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return
	}
	cmd.Server.GET("/users", func(c *gin.Context) {
		handlerUser.GetUser(c, db)
	})

	cmd.Server.POST("/user", func(c *gin.Context) {
		handlerUser.PostUser(c, db)
	})
	defer db.Close()
	err = StartMain(cmd.Server)
	if err != nil {
		log.Fatalln("Неудачевя запуск сервера")
	}
}

func StartMain(server *gin.Engine) error {
	log.Println("Сервер запущен!")
	return server.Run(":8080")
}
