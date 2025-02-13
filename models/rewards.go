package models

import "time"

type Rewards struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cost        int       `json:"cost"`
	EarnedAt    time.Time `json:"earned_at"`
}
