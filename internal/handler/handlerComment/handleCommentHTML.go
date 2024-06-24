package handlerComment

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
)

func GETHandlePostCommentsHTML(c *gin.Context, db *sqlx.DB) {
	if c.Param("idPost") == "" {
		c.HTML(400, "400.html", gin.H{"Error": "idPost can't be empty"})
		return
	}

	comments, err := services.GetCommentsByPostIdFull(c.Param("idPost"), db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}
	for i, _ := range comments {
		comments[i].DateCreatedFormat = comments[i].DateCreated.Format("2006-01-02 15:04:05")
	}

	var fullPostAndComments models.FullPost

	post, err := services.GetPostById(c.Param("idPost"), db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	user, err := services.GetUserById(db, post.AuthorId)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	fullPostAndComments = models.FullPost{
		Id:                post.Id,
		Title:             post.Title,
		Text:              post.Text,
		DateCreated:       post.DateCreated,
		DateCreatedFormat: post.DateCreated.Format("2006-01-02 15:04:05"),
		AuthorId:          post.AuthorId,
		AuthorName:        user.NickName,
		Comments:          comments,
		CommentsCount:     len(comments),
	}

	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		log.Println("Пользователь не авторизован или сессия истекла fgdfg")
		userID = ""
	}

	var userName string
	if userID != "" {
		user, err = services.GetUserById(db, userID.(string))
		if err != nil {
			c.HTML(400, "400.html", gin.H{"Error": err.Error()})
			return
		}
		userName = user.NickName.String
	} else {
		userName = "Авторизуйтесь чтобы комментировать"
	}

	c.HTML(http.StatusOK, "PagePostComments.html", gin.H{"posts": fullPostAndComments, "userID": userID, "userName": userName})
	return
}
