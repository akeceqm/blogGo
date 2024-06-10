package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/internal/handler/handlerPost"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {

	// запросы ПОСТОВ

	server.GET("/posts", func(c *gin.Context) {
		handlerPost.GETHandlePost(c, db)
	})

	server.POST("/posts", func(c *gin.Context) {
		handlerPost.POSTHandlePost(c, db)
	})

	server.GET("/posts/:idAuthor", func(c *gin.Context) {
		handlerPost.GETHandlePostByAuthorId(c, db)
	})

	server.GET("/posts/date/:startDate/:endDate", func(c *gin.Context) {
		handlerPost.GETHandlePostByBetweenDate(c, db)
	})

	server.PUT("/posts", func(c *gin.Context) {
		handlerPost.PUTHandlePost(c, db)
	})

	server.DELETE("/posts/:id", func(c *gin.Context) {
		handlerPost.DELETEHandlePost(c, db)
	})
}
