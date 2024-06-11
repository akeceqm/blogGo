package handlerComment

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/database/models"
	"post/internal/services"
)

func GETHandleCommentsByAuthorId(c *gin.Context, db *sqlx.DB) {
	if c.Param("idAuthor") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "idAuthor can't be empty"})
		return
	}

	comments, err := services.GetCommentsByAuthorId(c.Param("idAuthor"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
	return
}

func GETHandleCommentsByPostId(c *gin.Context, db *sqlx.DB) {
	if c.Param("idPost") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "idPost can't be empty"})
		return
	}

	comments, err := services.GetCommentsByPostId(c.Param("idPost"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
	return
}

func POSTHandleCommentByPostId(c *gin.Context, db *sqlx.DB) {
	if c.Param("idPost") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "idPost can't be empty"})
		return
	}

	var newComment models.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if newComment.Text == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Text can't be empty"})
		return
	}
	if newComment.AuthorId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Author can't be empty"})
		return
	}

	err := services.CreateComment(newComment.Text, newComment.AuthorId, c.Param("idPost"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
	return
}

func DELETEHandleCommentById(c *gin.Context, db *sqlx.DB) {
	if c.Param("idComment") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "idComment can't be empty"})
		return
	}

	err := services.DeleteComment(c.Param("idComment"), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
	return
}

func PUTHandleCommentById(c *gin.Context, db *sqlx.DB) {
	var newComment models.Comment

	if c.Param("idComment") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "idComment can't be empty"})
		return
	}

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if newComment.Text == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": "Text can't be empty"})
		return
	}

	err := services.UpdateComment(c.Param("idComment"), newComment.Text, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
	return

}
