package internal

import (
	"github.com/go-chi/chi/v5"
	"Weather-Forecast-API/internal/handlers"

)

func RegisterRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/weather", handlers.GetWeather)
		r.Post("/subscribe", handlers.Subscribe)
		r.Get("/confirm/{token}", handlers.Confirm)
		r.Get("/unsubscribe/{token}", handlers.Unsubscribe)
	})
}
