package success

import (
	"go-fullstack/internal/handlers/creation"
	"go-fullstack/internal/models/card"
	"go-fullstack/internal/rend"
	"go-fullstack/internal/storage"
	"net/http"
)

func New(renderer *rend.Renderer, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pages.creation.success.New"

		ctx := r.Context()

		data := creation.PageData{
			Title:     "Результат | Мотиватор",
			Submitted: false,
		}

		log := renderer.Logger()

		err := r.ParseForm()
		if err != nil {
			log.Error("card not parsed", "error", err)
		} else {
			card := card.Card{
				Title:   r.FormValue("title"),
				Content: r.FormValue("content"),
				Author:  r.FormValue("author"),
			}

			err = storage.CreateCard(ctx, &card)
			if err != nil {
				log.Error("card not created", "error", err)
			}

			data.Submitted = true
		}

		renderer.Render(w, "templates/pages/creation/success.html", data)
	}
}
