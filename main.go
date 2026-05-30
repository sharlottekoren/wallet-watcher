package main

import (
	"log"
	"net/http"

	"github.com/sharlottekoren/wallet-watcher/internal/handlers"
	"github.com/sharlottekoren/wallet-watcher/internal/services"
)

func main() {
	// Initialize the TransactionService
	txService := services.NewTransactionService()
	// Create a new TransactionHandler with the TransactionService
	txHandler := handlers.NewTransactionHandler(txService)

	// Set up the HTTP route for transactions and associate it with the TransactionHandler
	http.HandleFunc("/transactions", txHandler.HandleTransactions)

	// Start the HTTP server on port 8080 and log any errors that occur
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}