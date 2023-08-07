package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
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

var db *sql.DB

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

	if len(req.URL) < 8 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(req.URL) // Генерация сокращенного URL-адреса

	_, err := db.Exec("INSERT INTO short_urls (short_url, long_url) VALUES ($1, $2)", shortURL, req.URL)
	if err != nil {
		http.Error(w, "Failed to save short URL", http.StatusInternalServerError)
		fmt.Println("Failed to save short URL, reason:", err)
		return
	}

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
	var longURL string
	err := db.QueryRow("SELECT long_url FROM short_urls WHERE short_url = $1", shortURL).Scan(&longURL)
	if err != nil {
		http.NotFound(w, r) // Если короткий URL-адрес не найден, возвращается ошибка 404
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound) // Перенаправление на длинный URL-адрес
}

func generateShortURL(longURL string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(longURL))      // Кодирование длинного URL-адреса в base64
	path := strings.ToLower(encoded[len(encoded)-10 : len(encoded)-2]) // Формирование короткого URL-адреса
	return path

}

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=url_shortener sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/short", shortHandler)

	log.Printf("Server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
