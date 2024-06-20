package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"
	"post/internal/services"
)

func InitRoutesHTML(server *gin.Engine, db *sqlx.DB) {
	authMiddleware := AuthMiddleware(db)

	// Маршруты без авторизации
	server.GET("/authorization", func(c *gin.Context) {
		c.HTML(200, "authorization.html", gin.H{})
	})
	server.GET("/registration", func(c *gin.Context) {
		c.HTML(200, "registration.html", gin.H{})
	})

	// Применяем middleware авторизации
	server.Use(authMiddleware)

	// Маршруты с авторизацией
	server.GET("/", func(c *gin.Context) {
		handlerIndex(db, c)
	})
	server.GET("/profileUser", func(c *gin.Context) {
		c.HTML(200, "profileUser.html", gin.H{})
	})
	server.GET("/profileUser/:userId", func(c *gin.Context) {
		c.HTML(200, "profileUser.html", gin.H{})
	})
	server.GET("/changeProfile/:userId", func(c *gin.Context) {
		c.HTML(200, "changeProfile.html", gin.H{})
	})
	server.GET("/changeProfile", func(c *gin.Context) {
		c.HTML(200, "changeProfile.html", gin.H{})
	})
	server.GET("/h/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.GETHandlePostCommentsHTML(c, db)
	})
	server.GET("/h/:countPage", func(c *gin.Context) {
		handlerPost.GETHandlePostsHTML(c, db)
	})
}

func handlerIndex(db *sqlx.DB, c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		log.Println("Пользователь не авторизован или сессия истекла")
		handlerIndexNoAuthorization(c)
		return
	}

	// Проверка авторизации
	isAuthorized, err := services.IsUserAuthorized(db, userID.(string))
	if err != nil {
		log.Println("Ошибка проверки авторизации:", err)
		handlerIndexNoAuthorization(c)
		return
	}

	if isAuthorized {
		handlerIndexAuthorization(c)
	} else {
		handlerIndexNoAuthorization(c)
	}
}

func handlerIndexNoAuthorization(c *gin.Context) {
	log.Println("Rendering PageMainNoAuthorization.html")
	c.HTML(200, "PageMainNoAuthorization.html", nil)
}

func handlerIndexAuthorization(c *gin.Context) {
	log.Println("Rendering PagePostComments.html")
	c.HTML(200, "PagePostComments.html", nil)
}

func AuthMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			handlerIndexNoAuthorization(c)
			c.Abort()
			return
		}

		session, err := services.GetSessionByID(db, sessionID)
		if err != nil || session.UserID == "" {
			handlerIndexNoAuthorization(c)
			c.Abort()
			return
		}

		c.Set("userID", session.UserID) // Установка userID в контекст Gin для авторизованных пользователей
		c.Next()
	}
}
