package create

import (
	"go-fullstack/internal/models/card"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title     string
	Submitted bool
}

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create.New"
		const internalError = "Ошибка на стороне сервера."

		tmpl, err := template.ParseFiles("templates/layout.html", "templates/create.html")

		if err != nil {
			http.Error(w, internalError, http.StatusInternalServerError)
			log.Printf("%s: %s", op, err.Error())
			return
		}

		data := PageData{
			Title:     "Создать | Мотиватор",
			Submitted: false,
		}

		if r.Method == http.MethodPost {
			r.ParseForm()
			card := card.Card{
				Title:   r.FormValue("title"),
				Content: r.FormValue("content"),
				Author:  r.FormValue("author"),
			}
			data.Submitted = true
			log.Println(card)
		}

		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, internalError, http.StatusInternalServerError)
			log.Printf("%s: %s", op, err.Error())
		}
	}
}
