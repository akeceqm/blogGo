package handlerPost

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
	"time"
)

func POSTHandlePost(c *gin.Context, db *sqlx.DB) {
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

	createdPost, err := services.CreatePost(newPost.Title, newPost.Text, newPost.AuthorId, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Post": createdPost, "Success": "Post created"})
	return
}

func GETHandlePost(c *gin.Context, db *sqlx.DB) {
	post, err := services.GetPost(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
	return
}

func GETHandlePostByAuthorId(c *gin.Context, db *sqlx.DB) {
	if c.Param("idAuthor") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "idAuthor can't be empty"})
		return
	}

	post, err := services.GetPostsByAuthorId(c.Param("idAuthor"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
	return
}

func GETHandlePostByBetweenDate(c *gin.Context, db *sqlx.DB) {
	var posts []models.Post
	var startDate time.Time
	var endDate time.Time

	if c.Param("startDate") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "startDate can't be empty"})
		return
	}
	startDate, err := time.Parse(customFormat, c.Param("startDate"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	if c.Param("endDate") == "" {
		endDate = time.Now()
		return
	} else {
		endDate, err = time.Parse(customFormat, c.Param("endDate"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
			return
		}
	}

	posts, err = services.GetPostByBetweenDate(startDate, endDate, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
	return
}

func PUTHandlePost(c *gin.Context, db *sqlx.DB) {
	var updatedPost models.Post

	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := services.UpdatePost(updatedPost.Id, updatedPost.Title, updatedPost.Text, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Post updated"})
	return
}

func DELETEHandlePost(c *gin.Context, db *sqlx.DB) {

	err := services.DeletePost(c.Param("id"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Post deleted"})
	return
}
