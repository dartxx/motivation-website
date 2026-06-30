package main

import (
	"context"
	"go-fullstack/internal/config"
	"go-fullstack/internal/handlers/create"
	"go-fullstack/internal/handlers/root"
	"go-fullstack/internal/storage"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := config.MustLoad()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	storage, err := storage.New(ctx, cfg.ConnectionString)
	if err != nil {
		log.Error("Failed to init database", "error", err)
		os.Exit(1)
	}
	defer storage.Close()

	log.Info("Сервер запущен на - http://localhost:8080")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", root.New())
	http.HandleFunc("/create", create.New())

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Server start failed", "error", err)
	}
}
