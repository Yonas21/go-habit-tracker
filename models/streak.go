package models

import "time"

type Streak struct {
	ID            int       `json:"id"`
	HabitID       int       `json:"habit_id"`
	CurrentStreak int       `json:"current_streak"`
	LongestStreak int       `json:"longest_streak"`
	LastCompleted time.Time `json:"last_completed"`
}
