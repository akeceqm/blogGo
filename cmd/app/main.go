package main

import (
	"log"
	"net/http"

	"post/cmd"
	"post/internal/database"
	"post/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	cmd.Server = gin.Default()
	cmd.Server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	db, err := database.InitDb(database.ConnectionString)
	if err != nil {
		log.Fatalln("Неудачная попытка соединения с БД")
		return
	}
	defer db.Close()

	cmd.Server.Static("/assets/", "src/")
	cmd.Server.LoadHTMLGlob("src/html/*.html")
	handler.InitRoutes(cmd.Server, db)
	handler.InitRoutesHTML(cmd.Server, db)
	err = StartMain(cmd.Server)
	if err != nil {
		log.Fatalln("Неудачный запуск сервера")
	}
}

func StartMain(server *gin.Engine) error {
	log.Println("Сервер запущен")
	return server.Run(":8080")
}
