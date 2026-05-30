package main

import (
	"log"
	"net/http"

	"github.com/sharlottekoren/wallet-watcher/internal/handlers"
	"github.com/sharlottekoren/wallet-watcher/internal/services"

	"github.com/rs/cors"
)

func main() {
	// Initialize the transaction service and handler
	transactionService := services.NewTransactionService()
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Set up the HTTP server and routes
	mux := http.NewServeMux()
	mux.HandleFunc("/transactions", transactionHandler.HandleTransactions)

	// Set up CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	})

	// Wrap the HTTP handler with CORS middleware
	handler := c.Handler(mux)

	// Start the HTTP server with CORS enabled
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}