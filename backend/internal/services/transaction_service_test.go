package services

import (
	"testing"

	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
)

type mockTransactionRepository struct {
	transactions []models.Transaction
}

func (m *mockTransactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	m.transactions = append(m.transactions, transaction)
	return transaction, nil
}

func (m *mockTransactionRepository) GetAll() ([]models.Transaction, error) {
	transactions := make([]models.Transaction, len(m.transactions))
	copy(transactions, m.transactions)
	return transactions, nil
}

func Test_CreateTransaction(t *testing.T) {
	repo := &mockTransactionRepository{}
	service := NewTransactionService(repo)
	transaction := models.Transaction{
		Amount:          100,
		Description:     "Test transaction",
		TransactionType: "income",
	}

	createdTransaction, err := service.CreateTransaction(transaction)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if createdTransaction.ID == "" {
		t.Errorf("Expected transaction ID to be set, got empty string")
	}
	if createdTransaction.UserID != "user123" {
		t.Errorf("Expected UserID to be 'user123', got %s", createdTransaction.UserID)
	}
	if createdTransaction.Amount != transaction.Amount {
		t.Errorf("Expected Amount to be %d, got %d", transaction.Amount, createdTransaction.Amount)
	}
	if createdTransaction.Description != transaction.Description {
		t.Errorf("Expected Description to be '%s', got '%s'", transaction.Description, createdTransaction.Description)
	}
}

func Test_CreateTransaction_InvalidAmount(t *testing.T) {
	repo := &mockTransactionRepository{}
	service := NewTransactionService(repo)
	transaction := models.Transaction{
		Amount:          -50,
		Description:     "Invalid transaction",
		TransactionType: "income",
	}

	_, err := service.CreateTransaction(transaction)
	if err == nil {
		t.Fatal("Expected error for invalid amount, got nil")
	}
	if err.Error() != "amount must be greater than zero" {
		t.Errorf("Expected error message 'amount must be greater than zero', got '%s'", err.Error())
	}
}

func Test_CreateTransaction_EmptyDescription(t *testing.T) {
	repo := &mockTransactionRepository{}
	service := NewTransactionService(repo)
	transaction := models.Transaction{
		Amount:          50,
		Description:     "",
		TransactionType: "income",
	}

	_, err := service.CreateTransaction(transaction)
	if err == nil {
		t.Fatal("Expected error for empty description, got nil")
	}
	if err.Error() != "description cannot be empty" {
		t.Errorf("Expected error message 'description cannot be empty', got '%s'", err.Error())
	}
}

func Test_GetTransactions(t *testing.T) {
	// Arrange: Seed data directly into the mock repo slice
	repo := &mockTransactionRepository{
		transactions: []models.Transaction{
			{
				ID:              "tx-1",
				UserID:          "user123",
				Amount:          100,
				Description:     "Test transaction 1",
				TransactionType: "income",
			},
			{
				ID:              "tx-2",
				UserID:          "user123",
				Amount:          200,
				Description:     "Test transaction 2",
				TransactionType: "expense",
			},
		},
	}
	service := NewTransactionService(repo)

	// Act: Retrieve transactions directly from the service
	transactions, err := service.GetTransactions()

	// Assert: Verify the results match our expectations
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(transactions) != 2 {
		t.Fatalf("Expected 2 transactions, got %d", len(transactions))
	}
	if transactions[0].Description != "Test transaction 1" {
		t.Errorf("Expected first transaction description to be 'Test transaction 1', got '%s'", transactions[0].Description)
	}
	if transactions[1].Description != "Test transaction 2" {
		t.Errorf("Expected second transaction description to be 'Test transaction 2', got '%s'", transactions[1].Description)
	}
}