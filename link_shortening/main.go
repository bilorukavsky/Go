package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"` // Длинный URL-адрес, который нужно сократить
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"` // Сокращенный URL-адрес
}

var shortURLs map[string]string // Карта для хранения соответствий коротких и длинных URL-адресов

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req) // Разбор JSON-тела запроса
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(req.URL) // Генерация сокращенного URL-адреса
	shortURLs[shortURL] = req.URL         // Сохранение соответствия короткого и длинного URL-адресов

	resp := ShortenResponse{
		ShortURL: shortURL,
	}

	jsonResp, err := json.Marshal(resp) // Кодирование ответа в формат JSON
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp) // Отправка JSON-ответа
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[len("/redirect/"):]
	longURL, ok := shortURLs[shortURL] // Поиск соответствующего длинного URL-адреса по короткому URL-адресу
	if !ok {
		http.NotFound(w, r) // Если короткий URL-адрес не найден, возвращается ошибка 404
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound) // Перенаправление на длинный URL-адрес
}

func generateShortURL(longURL string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(longURL)) // Кодирование длинного URL-адреса в base64
	return encoded[:8]                                            // Возвращение первых 8 символов кодированной строки в качестве короткого URL-адреса
}

func main() {
	shortURLs = make(map[string]string) // Инициализация карты

	// Установка обработчиков для HTTP-запросов
	http.HandleFunc("/shorten", shortenHandler)    // Обработчик для сокращения URL-адреса
	http.HandleFunc("/redirect/", redirectHandler) // Обработчик для перенаправления

	log.Fatal(http.ListenAndServe(":8080", nil)) // Запуск сервера на порту 8080
}
