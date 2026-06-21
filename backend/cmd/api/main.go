package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sharlottekoren/wallet-watcher/backend/internal/handlers"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/repository"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/services"
	_ "modernc.org/sqlite"
)

func main() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the repository and service
	repo, err := repository.NewSQLiteRepository(db)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	service := services.NewTransactionService(repo)

	// Initialize the handler
	handler := handlers.NewTransactionHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTransactionHandler(w, r)
		case http.MethodGet:
			handler.GetTransactionsHandler(w, r)
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

	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
