package session

import (
	"net/http"
	"post/internal/middlewares"
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
