package repositories

import (
	"context"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

type OrderRepo interface {
	GetOrders(ctx context.Context) ([]entities.Order, error)
	GetNotProcessedOrders() ([]entities.Order, error)
	GetOrderByNumber(ctx context.Context, number string) (entities.Order, error)
	CheckIfOrderWasAddedByAnotherUser(ctx context.Context, number string) (bool, error)
	UpdateOrderStatus(orderID int, status int) error
	CreateOrder(ctx context.Context, number string) (entities.Order, error)
}
