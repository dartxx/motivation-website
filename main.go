package main

import (
	"context"
	"go-fullstack/internal/config"
	"go-fullstack/internal/handlers/root"
	"go-fullstack/internal/storage"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := config.MustLoad()

	storage, err := storage.New(ctx, cfg.ConnectionString)
	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
		os.Exit(1)
	}
	defer storage.Close()

	log.Println("Сервер запущен на - http://localhost:8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", root.New())

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
