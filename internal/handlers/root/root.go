package root

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.root.New"
		const internalError = "Ошибка на стороне сервера."

		tmpl, err := template.ParseFiles("templates/layout.html", "templates/root.html")

		if err != nil {
			http.Error(w, internalError, http.StatusInternalServerError)
			log.Printf("%s: %s", op, err.Error())
			return
		}

		data := PageData{
			Title: "Какой-то title",
		}

		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, internalError, http.StatusInternalServerError)
			log.Printf("%s: %s", op, err.Error())
		}
	}
}
