package repositories

import (
	"context"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

type TransRepo interface {
	GetBalance(ctx context.Context) (entities.Balance, error)
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)
	CreateTransaction(ctx context.Context, transactionDto dto.Transaction) error
}
