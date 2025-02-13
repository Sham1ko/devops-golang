package handlers

import (
	"devops-go/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок JSON
	w.Header().Set("Content-Type", "application/json")
	// Код ответа 200 OK
	w.WriteHeader(http.StatusOK)

	// Создаём JSON-ответ
	response := Response{Message: services.Hello()}

	// Кодируем JSON и отправляем в response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
	}

}
