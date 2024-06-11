package handler

import (
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"post/internal/handler/handlerUser"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {
	// Работа с юзером
	server.GET("/users", func(c *gin.Context) {
		handlerUser.GetHandleUsers(c, db)
	})
	server.POST("/users/authorization", func(c *gin.Context) {
		handlerUser.PostHandleAuthorizationUser(c, db)
	})
	server.POST("/users/registration", func(c *gin.Context) {
		handlerUser.PostHandleRegistrationUser(c, db)
	})

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

	// Запросы КОММЕНТАРИЕВ
	server.GET("/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.GETHandleCommentsByPostId(c, db)
	})

	server.GET("/user/:idAuthor/comments/", func(c *gin.Context) {
		handlerComment.GETHandleCommentsByAuthorId(c, db)
	})

	server.POST("/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.POSTHandleCommentByPostId(c, db)
	})

	server.PUT("/post/:idPost/comments/:idComment", func(c *gin.Context) {
		handlerComment.PUTHandleCommentById(c, db)
	})

	server.DELETE("/post/:idPost/comments/:idComment", func(c *gin.Context) {
		handlerComment.DELETEHandleCommentById(c, db)
	})

}
