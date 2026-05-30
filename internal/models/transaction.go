package models

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}