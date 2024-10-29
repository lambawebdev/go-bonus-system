package repositories

import (
	"context"
	"database/sql"
	"math"
	"time"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/middleware"
)

const GetWithdrawals = `
	SELECT number, sum, processed_at
	FROM withdrawals
	WHERE user_id = $1	
	`

const CreateWithdrawal = `
    INSERT INTO withdrawals (number, sum, processed_at, user_id)
	VALUES ($1, $2, $3, $4)
    `

type WithdrawalRepository struct {
	database *sql.DB
}

func NewWithdrawalRepository(db *sql.DB) *WithdrawalRepository {
	return &WithdrawalRepository{
		database: db,
	}
}

func (repository *WithdrawalRepository) GetWithdrawals(ctx context.Context) ([]entities.Withdrawal, error) {
	rows, err := repository.database.Query(GetWithdrawals, ctx.Value(&middleware.UserIDkey))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var withdrawals []entities.Withdrawal

	for rows.Next() {
		var amount float64
		var withdraw entities.Withdrawal

		if err := rows.Scan(&withdraw.Order, &amount, &withdraw.ProcessedAt); err != nil {
			return withdrawals, err
		}

		withdraw.Sum = math.Abs(amount) / 10000
		withdrawals = append(withdrawals, withdraw)
	}

	if err = rows.Err(); err != nil {
		return withdrawals, err
	}

	return withdrawals, nil
}

func (repository *WithdrawalRepository) CreateWithdrawal(ctx context.Context, transactionDto dto.Transaction) error {
	//Приводим к целочисленному значению перед сохранением
	amountToInt := math.Floor((transactionDto.Amount) * 10000)

	_, err := repository.database.Exec(CreateWithdrawal, transactionDto.Number, amountToInt, time.Now(), ctx.Value(&middleware.UserIDkey))

	if err != nil {
		return err
	}

	return nil
}
