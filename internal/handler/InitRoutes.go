package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/internal/handler/handlerPost"
	"post/internal/handler/handlerUser"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {
	// Рабоат с юзером
	server.GET("/users", func(c *gin.Context) {
		handlerUser.GetHandleUsers(c, db)
	})
	server.POST("/users/authorization", func(c *gin.Context) {
		handlerUser.PostHandleAuthorizationUser(c, db)
	})
	server.POST("/users/registration", func(c *gin.Context) {
		handlerUser.PostHandleRegistationUser(c, db)
	})

	// Рабоат с постами
	server.GET("/posts", func(c *gin.Context) {
		handlerPost.GETHandlePost(c, db)
	})

	server.POST("/posts", func(c *gin.Context) {
		handlerPost.POSTHandlePost(c, db)
	})
}
