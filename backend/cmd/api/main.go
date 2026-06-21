package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import driver for side-effects
	"github.com/sharlottekoren/wallet-watcher/backend/internal/handlers"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/repository"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/services"
)

func main() {
	// 1. Open connection to the SQLite database file
	db, err := sql.Open("sqlite3", "./wallet.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// 2. Setup database tables if they don't exist yet
	createTables(db)

	// 3. Initialize layers using Dependency Injection
	transactionRepo := repository.NewSQLiteTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
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

// Helper function to initialize our database on startup
func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		category_id TEXT NOT NULL,
		amount INTEGER NOT NULL,
		transaction_type TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}
