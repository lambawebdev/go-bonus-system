package repositories

import (
	"context"
	"database/sql"
	"math"
	"time"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	blackboxservice "github.com/lambawebdev/go-bonus-system/internal/gophemart/services/blackBoxService"
)

const GET_BALANCE = `
	SELECT COALESCE(SUM(transactions.amount), 0) - COALESCE(SUM(withdrawals.sum), 0) as current,
	       COALESCE(SUM(withdrawals.sum), 0) as withdrawn
	FROM transactions
	FULL JOIN withdrawals ON transactions.user_id = withdrawals.user_id
	WHERE transactions.user_id = $1	
	`

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

func (repository *TransactionRepository) GetBalance(ctx context.Context) (entities.Balance, error) {
	var balance entities.Balance

	if err := repository.database.QueryRow(GET_BALANCE, ctx.Value("user_id")).Scan(&balance.Current, &balance.Withdrawn); err != nil {
		if err == sql.ErrNoRows {
			return balance, err
		}
		return balance, err
	}

	balance.Current = balance.Current / 10000
	balance.Withdrawn = math.Abs(balance.Withdrawn / 10000)

	return balance, nil
}

func (repository *TransactionRepository) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	rows, err := repository.database.Query(GET_TRANSACTIONS, ctx.Value("user_id"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entities.Transaction

	for rows.Next() {
		var amount float64
		var transaction entities.Transaction

		if err := rows.Scan(&transaction.Number, &amount, &transaction.ProcessedAt); err != nil {
			return transactions, err
		}

		transaction.Amount = math.Abs(amount) / 10000
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) CreateTransaction(userId int, orderAccrual blackboxservice.OrderAccrual) error {
	//Приводим к целочисленному значению перед сохранением
	amountToInt := math.Floor((orderAccrual.Accrual) * 10000)

	_, err := repository.database.Exec(CREATE_TRANSACTION, amountToInt, userId, orderAccrual.OrderId, time.Now())

	if err != nil {
		return err
	}

	return nil
}
