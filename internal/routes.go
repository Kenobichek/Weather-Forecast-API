package internal

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/weather", GetWeather)
		r.Post("/subscribe", Subscribe)
		r.Get("/confirm/{token}", Confirm)
		r.Get("/unsubscribe/{token}", Unsubscribe)
	})
}
