package handlerUser

import (
	"net/http"
	"post/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func PostHandleRegistrationUser(c *gin.Context, db *sqlx.DB) {

	var emailRequest struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
		return
	}
	user, err := services.PostUser(db, emailRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"login": user.Login, "password": user.PasswordHash})

}
