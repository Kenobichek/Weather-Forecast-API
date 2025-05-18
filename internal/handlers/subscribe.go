package handlers

import (
	"encoding/json"
	"net/http"

	"Weather-Forecast-API/internal/models"
	"Weather-Forecast-API/internal/repository"
	"Weather-Forecast-API/internal/utilities"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	channel_type := r.FormValue("channel_type")
	channel_value := r.FormValue("channel_value")
	city := r.FormValue("city")
	frequency := r.FormValue("frequency")

	if channel_type == "" || channel_value == "" || city == "" || frequency == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if !utilities.IsValidChannel(channel_type) {
		http.Error(w, "Unsupported channel_type", http.StatusBadRequest)
		return
	}

	frequency_minutes, err := utilities.ConvertFrequency(frequency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sub := &models.Subscription{
		ChannelType:      channel_type,
		ChannelValue:     channel_value,
		City:             city,
		FrequencyMinutes: frequency_minutes,
	}

	if err := repository.CreateSubscription(sub); err != nil {
		http.Error(w, "Already subscribed or DB error", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Subscription successful. Confirmation sent."})
}
