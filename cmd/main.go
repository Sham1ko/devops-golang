package main

import (
	"devops-go/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)

	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
