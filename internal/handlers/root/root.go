package root

import (
	"go-fullstack/internal/rend"
	"net/http"
)

type PageData struct {
	Title string
}

func New(renderer *rend.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pages.root.New"

		data := PageData{
			Title: "Главная | Мотиватор",
		}

		renderer.Render(w, "templates/pages/root/root.html", data)
	}
}
