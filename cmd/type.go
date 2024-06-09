package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	Server *gin.Engine
	Db     *sqlx.DB
	Err    error
)
