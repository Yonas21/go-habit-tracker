package models

import "time"

type User struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Password_hash string    `json:"password_hash"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
