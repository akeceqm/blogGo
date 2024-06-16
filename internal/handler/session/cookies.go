package session

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"post/internal/middlewares"
	"post/internal/services"
)

func CookiesHandler(w http.ResponseWriter, r *http.Request) {
	// Читаем куки из запроса
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		// Если куки не существует или пустое значение, создаем новую сессию
		sessionID, err := middlewares.GenerateSessionID()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Устанавливаем новую куку с идентификатором сессии
		cookie := http.Cookie{
			Name:  "session_id",
			Value: sessionID,
			Path:  "/",
			// Дополнительные параметры, например, Expires или MaxAge, можно добавить по необходимости
		}
		http.SetCookie(w, &cookie)

		// Отправляем ответ о том, что сессионные куки установлены
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Сессионные куки установлены"))
		return
	}

	// Если куки уже существуют, можно просто отправить ответ о том, что все в порядке
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Сессионные куки уже установлены"))
}

func GetHandleUserById(c *gin.Context, db *sqlx.DB) {
	userId := c.Param("userId")

	// Вызываем сервис для получения данных пользователя из базы данных
	user, err := services.GetUserById(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "внутренняя ошибка сервера"})
		return
	}

	// Возвращаем данные пользователя в ответе
	c.JSON(http.StatusOK, gin.H{
		"id":                user.Id,
		"nick_name":         user.NickName,
		"registration_data": user.DateRegistration,
		"description":       user.Description,
		// Добавьте другие поля пользователя по необходимости
	})
}
