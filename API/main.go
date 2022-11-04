package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/cors"

	"api/handler"
)

func main() {
	port := ":3000"
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", func(r chi.Router) {
		r.Get("/{from}-{limit}", handler.GetLastMails())
		r.Get("/search/{term}-{from}-{limit}", handler.Search())
	})

	http.ListenAndServe(port, r)

}
