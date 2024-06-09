package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {
	server.GET("/posts", func(c *gin.Context) {
		HandlePostGET(c, db)
	})

	server.POST("/posts", func(c *gin.Context) {
		HandlePostPOST(c, db)
	})
}
