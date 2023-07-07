package main

import (
	"log"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	longURL := "https://www.example.com/some/long/url"
	expectedShortURL := "aHR0cHM6"

	shortURL := generateShortURL(longURL)

	if shortURL != expectedShortURL {
		t.Errorf("Ожидался сокращенный URL-адрес %s, но получен %s", expectedShortURL, shortURL)
	}
	log.Println("Тест 1 пройден успешно")

	// Тест с другим длинным URL-адресом
	anotherLongURL := "https://www.example.com/another/long/url"
	anotherExpectedShortURL := "aHR0cHM6"

	anotherShortURL := generateShortURL(anotherLongURL)

	if anotherShortURL != anotherExpectedShortURL {
		t.Errorf("Ожидался сокращенный URL-адрес %s, но получен %s", anotherExpectedShortURL, anotherShortURL)
	}
	log.Println("Тест 2 пройден успешно")
}
