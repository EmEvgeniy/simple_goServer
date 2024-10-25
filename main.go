package main

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Main Handler")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	htmlContent, err := os.ReadFile("about.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		log.Printf("Ошибка чтения файла: %v", err)
		return
	}

	w.Write(htmlContent)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Printf("Сервер запущен на порту %v \n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
