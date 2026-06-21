package models

import "time"

type Transaction struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	Amount      float64   `json:"amount" db:"amount"`
	Category    string    `json:"category" db:"category"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
