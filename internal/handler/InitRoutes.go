package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/internal/handler/handlerPost"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {
	server.GET("/posts", func(c *gin.Context) {
		handlerPost.GETHandlePost(c, db)
	})

	server.POST("/posts", func(c *gin.Context) {
		handlerPost.POSTHandlePost(c, db)
	})
}
