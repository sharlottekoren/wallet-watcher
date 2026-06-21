package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/repository"
	"strings"
	"time"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
}

// NewTransactionService creates a new instance of TransactionService with an empty transaction list.
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: repo,
	}
}

// CreateTransaction creates a new transaction and adds it to the service's transaction list.
func (s *TransactionService) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if transaction.Amount <= 0 {
		return models.Transaction{}, errors.New("amount must be greater than zero")
	}

	if strings.TrimSpace(transaction.Description) == "" {
		return models.Transaction{}, errors.New("description cannot be empty")
	}

	if transaction.TransactionType != "income" && transaction.TransactionType != "expense" {
		return models.Transaction{}, errors.New("transaction type must be either 'income' or 'expense'")
	}

	transaction.ID = uuid.New().String()
	transaction.UserID = "user123"
	transaction.CreatedAt = time.Now()

	_, err := s.transactionRepo.Create(transaction)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// GetTransactions returns all transactions in the service.
func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	transactions, err := s.transactionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
