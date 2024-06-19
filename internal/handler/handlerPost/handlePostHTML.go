package handlerPost

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/services"
	"strconv"
)

func GETHandlePostsHTML(c *gin.Context, db *sqlx.DB) {
	PageCount, err := strconv.Atoi(c.Param("countPage"))
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	fullPosts, err := services.GetFormFullPosts(PageCount, db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "PagePostsAll.html", gin.H{"posts": fullPosts})
	return
}
