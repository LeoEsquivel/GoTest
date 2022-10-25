package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"api/handler"
)

func main() {
	port := ":3000"
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/{from}-{limit}", handler.GetLastMails())
		r.Get("/search/{term}-{from}-{limit}", handler.Search())
	})

	http.ListenAndServe(port, r)

}
