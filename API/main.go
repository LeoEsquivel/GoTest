package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"api/handler"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/", handler.GetLastMails(0, 5))
	})

	http.ListenAndServe(":3000", r)
}
