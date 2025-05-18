package main

import (
	"log"
	"net/http"

	"Weather-Forecast-API/internal"
	"Weather-Forecast-API/internal/db"
	"Weather-Forecast-API/internal/scheduler"

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
	db.RunMigrations(db.DB)

	go scheduler.StartScheduler()

	r := chi.NewRouter()
	internal.RegisterRoutes(r)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
