package handlerUser

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"post/internal/database/models"
	"post/internal/services"
)

// Обработчик для получения данных пользователя по ID
func GetHandleUsers(c *gin.Context, db *sqlx.DB) {
	userId := c.Param("userId")

	user, err := services.GetUserById(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обработчик для обновления данных пользователя
func PUTHandleUser(c *gin.Context, db *sqlx.DB) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	userId := c.Param("userId")

	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	log.Printf("Updated user: %s, %s", updatedUser.NickName, updatedUser.Description)

	if updatedUser.Avatar != "" {
		// Decode base64 string to image
		avatarData, err := base64.StdEncoding.DecodeString(updatedUser.Avatar)
		if err != nil {
			log.Printf("Failed to decode base64 avatar: %v", err)
		} else {
			// Save image to file (e.g., in ./data/img/ with a unique name)
			avatarPath := SaveAvatarBase64(avatarData, userId)
			if avatarPath == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
				log.Println("Failed to save avatar")
				return
			}

			// Update user's avatar path in the structure
			updatedUser.Avatar = avatarPath
		}
	}

	// Update user in the database with new avatar path
	updatedUser, err := services.UpdateUser(db, userId, updatedUser.NickName, updatedUser.Description, updatedUser.Avatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		log.Printf("Failed to update user: %v", err)
		return
	}

	log.Println("Updated user avatar:", updatedUser.Avatar)

	// Return updated user as JSON
	c.JSON(http.StatusOK, updatedUser)
}

// Функция для сохранения файла аватара из base64 строки и возврата его пути
func SaveAvatarBase64(data []byte, userId string) string {
	uploadDir := "./data/imgAvatar/"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	filename := fmt.Sprintf("%s.jpg", userId)
	filePath := filepath.Join(uploadDir, filename)

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Printf("Failed to save avatar: %v", err)
		return ""
	}

	log.Printf("Avatar saved to %s", filePath)
	return "/data/imgAvatar/" + filename // Возвращаем путь для веб-сервера
}
