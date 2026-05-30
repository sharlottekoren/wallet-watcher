package services

import (
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
	"github.com/google/uuid"
	"time"
)

type TransactionService struct {
	transactions []models.Transaction
}

// NewTransactionService creates a new instance of TransactionService with an empty transaction list.
func NewTransactionService() *TransactionService {
	return &TransactionService{
		transactions: []models.Transaction{},
	}
}

// CreateTransaction creates a new transaction and adds it to the service's transaction list.
func (s *TransactionService) CreateTransaction(transaction models.Transaction) models.Transaction {
	transaction.ID = uuid.New().String()
	transaction.UserID = "user123"
	transaction.CreatedAt = time.Now()

	s.transactions = append(s.transactions, transaction)

	return transaction
}

// GetTransactions returns all transactions in the service.
func (s *TransactionService) GetTransactions() []models.Transaction {
	return s.transactions
}
