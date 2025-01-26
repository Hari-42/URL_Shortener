package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	URL string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			url := r.FormValue("url")

			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, PageData{URL: url})
			fmt.Printf("URL entered: %s\n", url)
		} else {
			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, nil)
		}
	})
	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
