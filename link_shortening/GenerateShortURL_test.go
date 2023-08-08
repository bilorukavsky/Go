package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
	tests := []struct {
		name        string
		longURL     string
		expectedURL string
	}{
		{
			name:        "Long URL 1",
			longURL:     "https://www.example.com/some/long/url",
			expectedURL: "5nl3vyba",
		},
		{
			name:        "URL = 8 char",
			longURL:     "test1.io",
			expectedURL: "dgvzddeu",
		},
		{
			name:        "URL < 8 char 'test.io'",
			longURL:     "test.io",
			expectedURL: "dgvzdc5",
		},
		{
			name:        "URL with single char ",
			longURL:     "t",
			expectedURL: "d",
		},
		{
			name:        "Long URL 2",
			longURL:     "https://www.example.com",
			expectedURL: "bszs5jb2",
		},
		{
			name:        "Long URL 3",
			longURL:     "https://www.example.com/another",
			expectedURL: "5vdghlcg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortURL := generateShortURL(tt.longURL)
			assert.Equal(t, tt.expectedURL, shortURL)
		})
	}
}
