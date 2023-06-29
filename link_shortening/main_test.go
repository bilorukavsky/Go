package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenHandler(t *testing.T) {
	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(shortenHandler))
	defer ts.Close()

	// Создаем тестовый запрос с длинным URL-адресом
	reqBody := ShortenRequest{URL: "https://example.com"}
	reqBodyJSON, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", ts.URL+"/shorten", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", resp.StatusCode)
	}

	// Проверяем формат ответа
	var respBody ShortenResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Fatal(err)
	}
	if respBody.ShortURL == "" {
		t.Error("Short URL is empty")
	}
}

func TestRedirectHandler(t *testing.T) {
	// Создаем карту с соответствиями коротких и длинных URL-адресов
	shortURLs = map[string]string{
		"abc123": "https://example.com",
	}

	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(redirectHandler))
	defer ts.Close()

	// Создаем тестовый запрос с коротким URL-адресом
	req, _ := http.NewRequest("GET", ts.URL+"/redirect/abc123", nil)

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusFound {
		t.Errorf("Expected status 302, but got %d", resp.StatusCode)
	}

	// Проверяем заголовок Location
	location := resp.Header.Get("Location")
	if location != "https://example.com" {
		t.Errorf("Expected redirect location to be 'https://example.com', but got '%s'", location)
	}
}
