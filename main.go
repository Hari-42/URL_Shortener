package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL-Shortener")
	})

	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
