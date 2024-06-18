package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"post/cmd"
	"post/internal/database/models"
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"
	"post/internal/services"
)

func InitRoutesHTML(server *gin.Engine, db *sqlx.DB) {

	cmd.Server.GET("/", func(c *gin.Context) {
		handlerIndex(c, db)
	})
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

func handlerIndex(c *gin.Context, db *sqlx.DB) {
	post, err := services.GetPostFull(db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	var fullPosts []models.FullPost
	for i := 0; i < len(post); i++ {
		comments, err := services.GetCommentsByPostId(post[i].Id, db)
		if err != nil {
			c.HTML(400, "400.html", gin.H{"Error": err.Error()})
			return
		}

		fullPosts = append(fullPosts, models.FullPost{
			Id:                post[i].Id,
			Title:             post[i].Title,
			Text:              post[i].Text,
			AuthorId:          post[i].AuthorId,
			DateCreatedFormat: post[i].DateCreated.Format("2006-01-02 15:04:05"),
			AuthorName:        post[i].AuthorName,
			Comments:          []models.FullComment{},
			CommentsCount:     len(comments),
		})
	}
	c.HTML(200, "PageMainNoAutorization.html", gin.H{"posts": fullPosts})
}
