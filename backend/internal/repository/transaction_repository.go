package repository

import (
	"database/sql"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
)

// TransactionRepository defines the interface for transaction data access operations.
type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionsByUserID(userID string) ([]*models.Transaction, error)
}

// SQLiteRepository implements the TransactionRepository interface for SQLite database.
type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository initialises the SQLite repository and creates the transactions table if it doesn't exist.
func NewSQLiteRepository(db *sql.DB) (*SQLiteRepository, error) {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
	id TEXT PRIMARY KEY,
	user_id TEXT NOT NULL,
	amount REAL NOT NULL,
	category TEXT NOT NULL,
	description TEXT,
	created_at DATETIME NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

// CreateTransaction inserts a new transaction into the database.
func (r *SQLiteRepository) CreateTransaction(transaction *models.Transaction) error {
	query := `INSERT INTO transactions (id, user_id, amount, category, description, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, transaction.ID, transaction.UserID, transaction.Amount, transaction.Category, transaction.Description, transaction.CreatedAt)
	return err
}

// GetTransactionsByUserID retrieves all transactions for a specific user from the database.
func (r *SQLiteRepository) GetTransactionsByUserID(userID string) ([]*models.Transaction, error) {
	query := `SELECT id, user_id, amount, category, description, created_at FROM transactions WHERE user_id = ?`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*models.Transaction, 0)

	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Category, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, &t)
	}
	return transactions, nil
}