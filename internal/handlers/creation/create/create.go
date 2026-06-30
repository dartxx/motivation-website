package create

import (
	"go-fullstack/internal/handlers/creation"
	"go-fullstack/internal/rend"
	"net/http"
)

func New(renderer *rend.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pages.creation.create.New"

		data := creation.PageData{
			Title: "Создать | Мотиватор",
		}

		renderer.Render(w, "templates/pages/creation/create.html", data)
	}
}
