package models

import "time"

type HabitLogs struct {
	ID          int       `json:"id"`
	HabitID     int       `json:"habit_id"`
	CompletedAt time.Time `json:"completed_at"`
}
