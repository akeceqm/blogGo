package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
)

func HandlePostPOST(c *gin.Context, db *sqlx.DB) {
	var newPost models.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if newPost.Title == "" || newPost.Text == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Title or text can't be empty"})
		return
	}
	if newPost.AuthorId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Author can't be empty"})
		return
	}

	err := services.CreatePost(newPost.Title, newPost.Text, newPost.AuthorId, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Post created"})
	return
}

func HandlePostGET(c *gin.Context, db *sqlx.DB) {
	post, err := services.GetPost(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
	return
}
