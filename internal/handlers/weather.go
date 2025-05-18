package handlers

import (
	"Weather-Forecast-API/internal/weather"
	"encoding/json"
	"net/http"
	"os"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing 'city' parameter", http.StatusBadRequest)
		return
	}

	provider := weather.OpenWeather{APIKey: os.Getenv("OPENWETHERMAP_API_KEY")}
	data, err := provider.GetWeather(city)
	
	if err != nil {
		if err.Error() == "city not found" {
			http.Error(w, "City not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to get weather: "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}