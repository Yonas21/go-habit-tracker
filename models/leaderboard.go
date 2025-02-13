package models

import "time"

type Leaderboard struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Score       int       `json:"score"`
	LastUpdated time.Time `json:"last_updated"`
}
