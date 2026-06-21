package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/repository"
	"time"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

// NewTransactionService creates a new instance of TransactionService with an empty transaction list.
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

// CreateTransaction creates a new transaction and adds it to the service's transaction list.
func (s *TransactionService) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	transaction.ID = uuid.New().String()
	transaction.UserID = "user123"
	transaction.CreatedAt = time.Now()

	if transaction.Amount <= 0 {
		return models.Transaction{}, errors.New("amount must be greater than zero")
	}

	if transaction.Description == "" {
		return models.Transaction{}, errors.New("description cannot be empty")
	}

	err := s.repo.CreateTransaction(&transaction)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// GetTransactions returns all transactions in the service.
func (s *TransactionService) GetTransactions() ([]*models.Transaction, error) {
	return s.repo.GetTransactionsByUserID("user123")
}
