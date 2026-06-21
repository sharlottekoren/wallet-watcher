package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

// NewTransactionHandler creates a new instance of TransactionHandler with the provided TransactionService.
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// CreateTransactionHandler handles the creation of a new transaction.
func (h *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request payload",
		})
		if err != nil {
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	createdTransaction, err := h.service.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		if err != nil {
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTransaction)
	if err != nil {
		http.Error(w, "Failed to encode created transaction", http.StatusInternalServerError)
	}
}

// GetTransactionsHandler handles the retrieval of all transactions.
func (h *TransactionHandler) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.service.GetTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to retrieve transactions from database",
		})
		if err != nil {
			http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
	}
}
