package repositories

import (
	"context"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	blackboxservice "github.com/lambawebdev/go-bonus-system/internal/gophemart/services/blackBoxService"
)

type TransRepo interface {
	GetBalance(ctx context.Context) (entities.Balance, error)
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)
	CreateTransaction(userIn int, orderAccrual blackboxservice.OrderAccrual) error
}
