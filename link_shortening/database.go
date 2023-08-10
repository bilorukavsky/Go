package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initializeDB() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=url_shortener sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func closeDB() {
	if db != nil {
		db.Close()
	}
}

func saveShortURL(shortURL, longURL string) error {
	_, err := db.Exec("INSERT INTO short_urls (short_url, long_url) VALUES ($1, $2)", shortURL, longURL)
	return err
}

func getLongURL(shortURL string) (string, error) {
	var longURL string
	err := db.QueryRow("SELECT long_url FROM short_urls WHERE short_url = $1", shortURL).Scan(&longURL)
	return longURL, err
}
