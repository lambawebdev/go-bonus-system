package repositories

import (
	"context"
	"database/sql"
	"math"
	"time"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

const GET_TRANSACTIONS = `
    SELECT orders.number, amount, processed_at FROM transactions
	LEFT JOIN orders ON transactions.order_id = orders.id
	WHERE transactions.user_id = $1
	ORDER BY processed_at DESC
    `
const CREATE_TRANSACTION = `
   INSERT INTO transactions (amount, user_id, order_id, processed_at)
   VALUES ($1, $2, $3, $4)
    `

type TransactionRepository struct {
	database *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		database: db,
	}
}

func (repository *TransactionRepository) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	rows, err := repository.database.Query(GET_TRANSACTIONS, ctx.Value("user_id"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entities.Transaction

	for rows.Next() {
		var amount int
		var transaction entities.Transaction

		if err := rows.Scan(&transaction.Number, &amount, &transaction.ProcessedAt); err != nil {
			return transactions, err
		}

		transaction.Amount = amount / 10000
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) CreateTransaction(ctx context.Context, transactionDto dto.Transaction) error {
	//Приводим к целочисленному значению перед сохранением
	amountToInt := math.Floor(transactionDto.Amount * 10000)

	_, err := repository.database.Exec(CREATE_TRANSACTION, amountToInt, ctx.Value("user_id"), transactionDto.OrderId, time.Now())

	if err != nil {
		return err
	}

	return nil
}
