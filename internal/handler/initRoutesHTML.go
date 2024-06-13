package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/cmd"
)

func InitRoutesHTML(server *gin.Engine, db *sqlx.DB) {

	cmd.Server.GET("/", handlerIndex)

	server.GET("/profileUser", func(c *gin.Context) {

		c.HTML(200, "profileUser.html", gin.H{})
	})

	server.GET("/registration", func(c *gin.Context) {
		c.HTML(200, "registration.html", gin.H{})
	})
	server.GET("/authorization", func(c *gin.Context) {
		c.HTML(200, "authorization.html", gin.H{})
	})
}
func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
