package repository

import (
	"Weather-Forecast-API/internal/db"
	"Weather-Forecast-API/internal/models"
	"errors"
	"strings"
	"time"
)

func CreateSubscription(s *models.Subscription) error {
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

func ConfirmByToken(token string) error {
	result, err := db.DB.Exec(`
		UPDATE subscriptions
		SET confirmed = TRUE
		WHERE token = $1`, token)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("not found")
	}

	return nil
}

func UnsubscribeByToken(token string) error {
	result, err := db.DB.Exec(`DELETE FROM subscriptions WHERE token = $1`, token)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("not found")
	}
	return nil
}
