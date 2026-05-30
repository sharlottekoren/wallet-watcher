package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/sharlottekoren/wallet-watcher/internal/services"
)

// TransactionHandler handles HTTP requests related to transactions.
type TransactionHandler struct {
	service *services.TransactionService
}

// NewTransactionHandler creates a new instance of TransactionHandler with the given TransactionService.
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// HandleTransactions processes HTTP requests for transactions. It supports GET to retrieve all transactions and POST to create a new transaction.
func (h *TransactionHandler) HandleTransactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	// Handle GET requests by returning the list of transactions as JSON
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.GetTransactions())

	// Handle POST requests by decoding the request body into a Transaction struct, creating a new transaction using the service, and returning it as JSON
	case http.MethodPost:
		var tx struct {
			Amount      float64 `json:"amount"`
			Category    string  `json:"category"`
			Description string  `json:"description"`
		}

		err := json.NewDecoder(r.Body).Decode(&tx)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		newTx := h.service.CreateTransaction(tx.Amount, tx.Category, tx.Description)
		json.NewEncoder(w).Encode(newTx)

	// Handle any other HTTP methods by responding with a "Method not allowed" error
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}