package repositories

import (
	"context"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

type WithdrawalRepo interface {
	GetWithdrawals(ctx context.Context) ([]entities.Withdrawal, error)
	CreateWithdrawal(ctx context.Context, transactionDto dto.Transaction) error
}
