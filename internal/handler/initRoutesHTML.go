package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/cmd"
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"
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
	server.GET("/h/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.GETHandlePostCommentsHTML(c, db)
	})
	server.GET("/h/:countPage", func(c *gin.Context) {
		handlerPost.GETHandlePostsHTML(c, db)
	})

}

func handlerIndex(c *gin.Context) {
	c.HTML(200, "PageMainNoAutorization.html", nil)
}
