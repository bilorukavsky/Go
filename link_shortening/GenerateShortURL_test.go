package main

import (
	"log"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	// Тест длинным URL-адресом 1
	longURL1 := "https://www.example.com/some/long/url"
	expectedShortURL1 := "5nl3vyba"

	shortURL1 := generateShortURL(longURL1)

	if shortURL1 != expectedShortURL1 {
		t.Errorf("Ожидался сокращенный URL-адрес %s, но получен %s", expectedShortURL1, shortURL1)
	}
	log.Println("Тест 1 пройден успешно")

	// Тест длинным URL-адресом 2
	LongURL2 := "https://www.example.com/another"
	expectedShortURL2 := "5vdghlcg"

	ShortURL2 := generateShortURL(LongURL2)

	if ShortURL2 != expectedShortURL2 {
		t.Errorf("Ожидался сокращенный URL-адрес %s, но получен %s", expectedShortURL2, ShortURL2)
	}
	log.Println("Тест 2 пройден успешно")

	// Тест длинным URL-адресом 3
	LongURL3 := "https://www.example.com"
	expectedShortURL3 := "bszs5jb2"

	ShortURL3 := generateShortURL(LongURL3)

	if ShortURL3 != expectedShortURL3 {
		t.Errorf("Ожидался сокращенный URL-адрес %s, но получен %s", expectedShortURL3, ShortURL3)
	}
	log.Println("Тест 3 пройден успешно")
}
