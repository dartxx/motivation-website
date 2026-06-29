package main

import (
	"go-fullstack/internal/handlers/root"
	"log"
	"net/http"
)

func main() {
	log.Println("Сервер запущен на - http://localhost:8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", root.New())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
