package repositories

import (
	"context"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

type TransRepo interface {
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)
	CreateTransaction(ctx context.Context, transactionDto dto.Transaction) error
}
