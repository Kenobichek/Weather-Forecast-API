package repository

import (
	"Weather-Forecast-API/internal/models"
	"Weather-Forecast-API/internal/db"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateSubscription(s *models.Subscription) error {
	s.Token = uuid.NewString()
	s.NextNotifiedAt = time.Now().Add(time.Duration(s.FrequencyMinutes) * time.Minute)

	_, err := db.DB.Exec(`
		INSERT INTO subscriptions 
		(channel_type, channel_value, city, frequency_minutes, token, next_notified_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		s.ChannelType, s.ChannelValue, s.City, 
		s.FrequencyMinutes, s.Token, s.NextNotifiedAt,
	)

	if err != nil && strings.Contains(err.Error(), "unique") {
		return errors.New("already subscribed")
	}
	return err
}
