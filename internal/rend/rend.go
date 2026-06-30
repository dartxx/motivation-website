package rend

import (
	"html/template"
	"log/slog"
	"net/http"
)

type Renderer struct {
	logger *slog.Logger
	layout string
}

func New(log *slog.Logger) *Renderer {
	return &Renderer{
		logger: log,
		layout: "templates/layout.html",
	}
}

func (r *Renderer) Render(w http.ResponseWriter, page string, data any) error {
	tmpl, err := template.ParseFiles(r.layout, page)
	if err != nil {
		r.logger.Error("failed to load template", "error", err, "page", page)
		http.Error(w, "Ошибка на стороне сервера.", http.StatusInternalServerError)
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		r.logger.Error("failed to execute template", "error", err)
		http.Error(w, "Ошибка на стороне сервера.", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (r *Renderer) Logger() *slog.Logger {
	return r.logger
}
