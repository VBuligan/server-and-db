package api

import "net/http"

// * Full API Handler initialization file

// Message * Вспомогательная структура для сообщений
type Message struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
