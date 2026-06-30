package root

import (
	"go-fullstack/internal/models/card"
	"go-fullstack/internal/rend"
	"go-fullstack/internal/storage"
	"net/http"
)

type PageData struct {
	Title string
	Cards []card.Card
}

func New(renderer *rend.Renderer, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pages.root.New"

		log := renderer.Logger()

		ctx := r.Context()

		cards, err := storage.GetAllCards(ctx)
		if err != nil {
			log.Error("failed to get cards", "error", err)
		}

		for i := range cards {
			cards[i].CreatedAtStr = cards[i].CreatedAt.Format("15:04 02-01")
		}

		data := PageData{
			Title: "Главная | Мотиватор",
			Cards: cards,
		}

		renderer.Render(w, "templates/pages/root/root.html", data)
	}
}
