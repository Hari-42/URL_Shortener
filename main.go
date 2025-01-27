package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type PageData struct {
	OriginalURL string
	ShortURL    string
}

var urlStore = make(map[string]string)
var mutex = &sync.Mutex{}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8

	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return "shorting.lio/" + string(shortURL)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			originalURL := r.FormValue("url")

			shortURL := generateShortURL()

			mutex.Lock()
			urlStore[shortURL] = originalURL
			mutex.Unlock()

			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, PageData{OriginalURL: originalURL, ShortURL: shortURL})
			fmt.Printf("Original URL: %s -> Shortened URL: %s\n", originalURL, shortURL)
		} else {
			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.HandleFunc("/shorting.lio/", func(w http.ResponseWriter, r *http.Request) {
		shortURL := "shorting.lio/" + r.URL.Path[len("/shorting.lio/"):]
		mutex.Lock()
		originalURL, exists := urlStore[shortURL]
		mutex.Unlock()

		if exists {
			http.Redirect(w, r, originalURL, http.StatusFound)
		} else {
			http.NotFound(w, r)
		}
	})

	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
