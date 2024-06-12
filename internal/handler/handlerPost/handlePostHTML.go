package handlerPost

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
	"strconv"
)

func GETHandlePostHTML(c *gin.Context, db *sqlx.DB) {
	post, err := services.GetPostFull(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	PageCount, err := strconv.Atoi(c.Param("countPage"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	var fullPosts []models.FullPost
	for i := (PageCount - 1) * 10; i < PageCount+10 && i < len(post); i++ {
		fullPosts = append(fullPosts, models.FullPost{
			Id:                post[i].Id,
			Title:             post[i].Title,
			Text:              post[i].Text,
			AuthorId:          post[i].AuthorId,
			DateCreatedFormat: post[i].DateCreated.Format("2006-01-02 15:04:05"),
			AuthorName:        post[i].AuthorName,
			Comments:          []models.FullComment{},
			CommentsCount:     0,
		})
	}

	c.HTML(http.StatusOK, "PagePostsAll.html", gin.H{"posts": fullPosts})
	return
}
