package models

import "time"

// Category represents a budget group like "Groceries" or "Rent"
type Category struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	TypeGroup string `json:"type_group"` // e.g., "Essentials", "Lifestyle"
}

// Transaction represents a single financial movement
type Transaction struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	CategoryID      string    `json:"category_id"`
	Amount          int64     `json:"amount"`           // Stored in lowest unit (e.g., pence)
	TransactionType string    `json:"transaction_type"` // "income" or "expense"
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
}
