package models

import "time"

type User struct {
	ID           int       `json:"user_id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	PasswordHash string    `json:"password_hash"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
}
