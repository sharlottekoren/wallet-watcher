package services

import (
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
	"testing"
)

func Test_CreateTransaction(t *testing.T) {
	service := NewTransactionService()
	transaction := models.Transaction{
		Amount:      100.0,
		Description: "Test transaction",
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
		t.Errorf("Expected Amount to be %f, got %f", transaction.Amount, createdTransaction.Amount)
	}
	if createdTransaction.Description != transaction.Description {
		t.Errorf("Expected Description to be '%s', got '%s'", transaction.Description, createdTransaction.Description)
	}
}

func Test_GetTransactions(t *testing.T) {
	service := NewTransactionService()

	transaction1 := models.Transaction{
		Amount:      100.0,
		Description: "Test transaction 1",
	}

	transaction2 := models.Transaction{
		Amount:      200.0,
		Description: "Test transaction 2",
	}

	_, err := service.CreateTransaction(transaction1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, err = service.CreateTransaction(transaction2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	transactions := service.GetTransactions()

	if len(transactions) != 2 {
		t.Fatalf("Expected 2 transactions, got %d", len(transactions))
	}
	if transactions[0].Description != transaction1.Description {
		t.Errorf("Expected first transaction description to be '%s', got '%s'", transaction1.Description, transactions[0].Description)
	}
	if transactions[1].Description != transaction2.Description {
		t.Errorf("Expected second transaction description to be '%s', got '%s'", transaction2.Description, transactions[1].Description)
	}
}

func Test_CreateTransaction_InvalidAmount(t *testing.T) {
	service := NewTransactionService()
	transaction := models.Transaction{
		Amount:      -50.0,
		Description: "Invalid transaction",
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
	service := NewTransactionService()
	transaction := models.Transaction{
		Amount:      50.0,
		Description: "",
	}

	_, err := service.CreateTransaction(transaction)
	if err == nil {
		t.Fatal("Expected error for empty description, got nil")
	}
	if err.Error() != "description cannot be empty" {
		t.Errorf("Expected error message 'description cannot be empty', got '%s'", err.Error())
	}
}
