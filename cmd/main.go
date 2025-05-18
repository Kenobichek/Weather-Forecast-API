package main

import (
	"log"
	"net/http"

	"Weather-Forecast-API/internal"
	"Weather-Forecast-API/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file.")
	}
}

func main() {
	db.Init()
	log.Println()

	r := chi.NewRouter()
	internal.RegisterRoutes(r)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
