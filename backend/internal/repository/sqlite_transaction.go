package repository

import (
	"database/sql"
	"github.com/sharlottekoren/wallet-watcher/backend/internal/models"
)

type SQLiteTransactionRepository struct {
	db *sql.DB
}

func NewSQLiteTransactionRepository(db *sql.DB) *SQLiteTransactionRepository {
	return &SQLiteTransactionRepository{db: db}
}

func (r *SQLiteTransactionRepository) Create(t models.Transaction) (models.Transaction, error) {
	query := `
		INSERT INTO transactions (id, user_id, category_id, amount, transaction_type, description, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	// Execute the raw SQL statement using the database driver
	_, err := r.db.Exec(query, t.ID, t.UserID, t.CategoryID, t.Amount, t.TransactionType, t.Description, t.CreatedAt)
	if err != nil {
		return models.Transaction{}, err
	}

	return t, nil
}

func (r *SQLiteTransactionRepository) GetAll() ([]models.Transaction, error) {
	query := `SELECT id, user_id, category_id, amount, transaction_type, description, created_at FROM transactions`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var t models.Transaction

		err := rows.Scan(&t.ID, &t.UserID, &t.CategoryID, &t.Amount, &t.TransactionType, &t.Description, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
