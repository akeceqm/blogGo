package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"post/cmd"
	"post/internal/database/models"
	"post/internal/handler/handlerComment"
	"post/internal/handler/handlerPost"
	"post/internal/handler/handlerUser"
	"post/internal/services"
)

func InitRoutesHTML(server *gin.Engine, db *sqlx.DB) {
	authMiddleware := AuthMiddleware(db)

	server.GET("/authorization", func(c *gin.Context) {
		c.HTML(200, "authorization.html", gin.H{})
	})

	server.GET("/registration", func(c *gin.Context) {
		c.HTML(200, "registration.html", gin.H{})
	})
	// Применяем middleware авторизации
	server.Use(authMiddleware)

	cmd.Server.GET("/", func(c *gin.Context) {
		handlerIndex(db, c)

	})

	server.GET("/profileUser", func(c *gin.Context) {
		handlerIndexProfileUser(c, db)
	})

	server.GET("/profileUser/:userId", func(c *gin.Context) {
		c.HTML(200, "profileUser.html", gin.H{})
	})
	server.GET("/changeProfile/:userId", func(c *gin.Context) {
		c.HTML(200, "changeProfile.html", gin.H{})
	})

	server.PUT("/changeProfile/:userId", func(c *gin.Context) {
		handlerUser.PUTHandleUser(c, db)

	})
	server.GET("/h/post/:idPost/comments", func(c *gin.Context) {
		handlerComment.GETHandlePostCommentsHTML(c, db)
	})
	server.GET("/h/:countPage", func(c *gin.Context) {
		handlerPost.GETHandlePostsHTML(c, db)
	})

	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
}

func handlerIndex(db *sqlx.DB, c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		handlerIndexNoAuthorization(c, db)
		return
	}

	// Проверка авторизации
	isAuthorized, err := services.IsUserAuthorized(db, userID.(string))
	if err != nil {
		log.Println("Ошибка проверки авторизации:", err)
		handlerIndexNoAuthorization(c, db)
		return
	}

	if isAuthorized {
		handlerIndexAuthorization(c, db)
	} else {
		handlerIndexNoAuthorization(c, db)
	}
}

func handlerIndexNoAuthorization(c *gin.Context, db *sqlx.DB) {
	log.Println("Rendering PageMainNoAuthorization.html")
	post, err := services.GetPostFull(db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	if len(post) == 0 {
		log.Println("No posts found")
		c.HTML(200, "PageMainNoAuthorization.html", gin.H{"posts": []models.FullPost{}})
		return
	}

	var fullPosts []models.FullPost
	for i := 0; i < 10 && i < len(post); i++ {
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
	c.HTML(200, "PageMainNoAuthorization.html", gin.H{"posts": fullPosts})
}

<<<<<<< Updated upstream
func handlerIndexAuthorization(c *gin.Context, db *sqlx.DB) {
	log.Println("Rendering PageMainYesAuthorization.html")
	post, err := services.GetPostFull(db)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	if len(post) == 0 {
		log.Println("No posts found")
	}

	var fullPosts []models.FullPost
	for i := 0; i < 10 && i < len(post); i++ {
		comments, err := services.GetCommentsByPostId(post[i].Id, db)
		if err != nil {
			continue
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

	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		log.Println("Пользователь не авторизован или сессия истекла fgdfg")
		userID = ""
	}

	userName, err := services.GetUserById(db, userID.(string))
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	c.HTML(200, "PageMainYesAuthorization.html", gin.H{"posts": fullPosts, "userName": userName})
}

func handlerIndexProfileUser(c *gin.Context, db *sqlx.DB) {
	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		log.Println("Пользователь не авторизован или сессия истекла fgdfg")
		userID = ""
	}

	post, err := services.GetPostFullByUserId(db, userID.(string))
	if err != nil {
		c.HTML(400, "400.html", gin.H{"Error": err.Error()})
		return
	}

	var fullPosts []models.FullPost
	for i := 0; i < 10 && i < len(post); i++ {
		comments, err := services.GetCommentsByPostId(post[i].Id, db)
		if err != nil {
			continue
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

	c.HTML(200, "profileUser.html", gin.H{"posts": fullPosts})
=======
// TODO сделать загрузку постов
func handlerIndexAuthorization(c *gin.Context) {
	log.Println("Rendering PagePostComments.html")
	c.HTML(200, "PageMainYesAuthorization.html", nil)

>>>>>>> Stashed changes
}

func AuthMiddleware(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			return
		}

		session, err := services.GetSessionByID(db, sessionID)
		if err != nil || session.UserID == "" {
			return
		}

		c.Set("userID", session.UserID) // Установка userID в контекст Gin для авторизованных пользователей
		c.Next()
	}
}
