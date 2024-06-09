package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"post/cmd"
	"post/internal/database"

func main() {
	cmd.Server = gin.Default()
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачевя попытка соединения с бд")
		return
	}
	defer db.Close()
	
	cmd.Server.GET("/users", func(c *gin.Context) {
		handlerUser.GetUsers(c, db)
	})

	cmd.Server.POST("/users", func(c *gin.Context) {
		handlerUser.PostUser(c, db)
	})
	err = StartMain(cmd.Server)
	if err != nil {
		log.Fatalln("Неудачевя попытка запуска сервера")
}

func StartMain(server *gin.Engine) error {
	log.Println("Сервер запущен")
	return server.Run(":8080")
}
