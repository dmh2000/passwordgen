package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	pwd "sqirvy.xyz/passwords/internal/pwd"
)

//go:embed static/*
var static embed.FS

type PasswordResponse struct {
	Passwords []Password `json:"passwords"`
}

type Password struct {
	Raw      string `json:"raw"`
	Formatted string `json:"formatted"`
}

func main() {
	// Parse template
	tmpl, err := template.ParseFS(static, "static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files
	http.Handle("/static/", http.FileServer(http.FS(static)))

	// Serve index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Handle password generation
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		length, _ := strconv.Atoi(r.URL.Query().Get("length"))
		if length < 24 {
			length = 24
		}

		count, _ := strconv.Atoi(r.URL.Query().Get("count"))
		if count < 1 {
			count = 1
		}

		symbols := r.URL.Query().Get("symbols") == "true"

		passwords := make([]Password, count)
		for i := 0; i < count; i++ {
			raw := pwd.GeneratePassword(length, symbols)
			passwords[i] = Password{
				Raw:      raw,
				Formatted: pwd.FormatPassword(raw),
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PasswordResponse{Passwords: passwords})
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
