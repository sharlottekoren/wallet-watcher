package services

import (
	"github.com/google/uuid"
	"github.com/sharlottekoren/wallet-watcher/internal/models"
)

// TransactionService provides methods to manage transactions.
type TransactionService struct {
	transactions []models.Transaction
}

// NewTransactionService creates a new instance of TransactionService.
func NewTransactionService() *TransactionService {
	return &TransactionService{
		transactions: []models.Transaction{},
	}
}

// CreateTransaction adds a new transaction to the service and returns it.
func (s *TransactionService) CreateTransaction(amount float64, category, description string) models.Transaction {
	tx := models.Transaction{
		ID:          uuid.New().String(),
		Amount:      amount,
		Category:    category,
		Description: description,
	}

	s.transactions = append(s.transactions, tx)
	return tx
}

// GetTransactions returns all transactions managed by the service.
func (s *TransactionService) GetTransactions() []models.Transaction {
	return s.transactions
}