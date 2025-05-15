package main


import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"Weather-Forecast-API/internal"
)

func main() {
	r := chi.NewRouter()
	internal.RegisterRoutes(r)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
