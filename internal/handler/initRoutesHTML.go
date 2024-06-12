package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"
)

func InitRoutesHTML(server *gin.Engine, db *sqlx.DB) {

	server.GET("/h/:countPage", func(c *gin.Context) {
		handlerPost.GETHandlePostHTML(c, db)
	})

	server.GET("/h/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.GETHandlePostCommentsHTML(c, db)
	})
}
