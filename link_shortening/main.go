package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "http://localhost:8080/"

type Origin struct {
	URL string `json:"url"` // Длинный URL-адрес, который нужно сократить
}

type Shortened struct {
	URL string `json:"short_url"` // Сокращенный URL-адрес
}

type Response struct {
	Data any `json:"data"`
}

var shortURLs map[string]string // Карта для хранения соответствий коротких и длинных URL-адресов

func isValidURL(longURL string) bool {
	parse, err := url.Parse(longURL)
	return err == nil && parse.Host != "" //&& parse.Scheme != ""
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Origin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // Разбор JSON-тела запроса
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if isValidURL(req.URL) != true {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(req.URL) // Генерация сокращенного URL-адреса
	shortURLs[shortURL] = req.URL         // Сохранение соответствия короткого и длинного URL-адресов

	resp := Response{
		Shortened{
			URL: baseURL + shortURL,
		},
	}

	jsonResp, err := json.Marshal(resp) // Кодирование ответа в формат JSON
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Отправка JSON-ответа. Обрабатываем ошибку, если не получилось отправить ответ
	if _, err := w.Write(jsonResp); err != nil {
		log.Println("Failed to write response:", err)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[len("/"):]
	longURL, ok := shortURLs[shortURL] // Поиск соответствующего длинного URL-адреса по короткому URL-адресу
	if !ok {
		http.NotFound(w, r) // Если короткий URL-адрес не найден, возвращается ошибка 404
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound) // Перенаправление на длинный URL-адрес
}

func generateShortURL(longURL string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(longURL)) // Кодирование длинного URL-адреса в base64
	if len(longURL) <= 8 {
		return strings.ToLower(encoded[:len(longURL)]) // Либо return longURL если возвращать без изменений
	} else {
		return strings.ToLower(encoded[len(encoded)-10 : len(encoded)-2])
	}
}

func main() {
	shortURLs = make(map[string]string) // Инициализация карты

	// Установка обработчиков для HTTP-запросов
	http.HandleFunc("/", redirectHandler)   // Обработчик для перенаправления
	http.HandleFunc("/short", shortHandler) // Обработчик для сокращения URL-адреса

	log.Fatal(http.ListenAndServe(":8080", nil)) // Запуск сервера на порту 8080

}
