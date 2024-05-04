// repository/transaction_repository.go
package repository

import (
	"context"
	"database/sql"

	"github.com/joshua468/banking-analytics/models"
)

type TransactionRepository interface {
	GetTransactions(ctx context.Context) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactions(ctx context.Context) ([]models.Transaction, error) {
	query := "SELECT id, account_id, amount, description FROM transactions"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Amount, &transaction.Description); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
