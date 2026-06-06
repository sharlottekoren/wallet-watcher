package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sharlottekoren/wallet-watcher/backend/internal/handlers"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/services"
)

func main() {
	transactionService := services.NewTransactionService()
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	mux := http.NewServeMux()
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.CreateTransactionHandler(w, r)
		case http.MethodGet:
			transactionHandler.GetTransactionsHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			err := json.NewEncoder(w).Encode(map[string]string{
				"error": "Method not allowed",
			})
			if err != nil {
				http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
			}
		}
	})

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
