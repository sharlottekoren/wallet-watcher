package repository

import "github.com/sharlottekoren/wallet-watcher/backend/internal/models"

// TransactionRepository defines the behavior our data layer must support
type TransactionRepository interface {
	Create(transaction models.Transaction) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
}
