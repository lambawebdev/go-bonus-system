package orderservice

import (
	"context"
	"fmt"
	"time"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	blackboxservice "github.com/lambawebdev/go-bonus-system/internal/gophemart/services/blackBoxService"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(orderRepo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (orderServices *OrderService) RunUpdateOrdersStatuses(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			orderServices.UpdateOrdersStatuses()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (orderService *OrderService) UpdateOrdersStatuses() {
	orders, err := orderService.orderRepo.GetNotProcessedOrders()

	if err != nil {
		fmt.Println(err)
	}

	for _, order := range orders {
		orderStatus, err := blackboxservice.GetOrderStatus(order.Number)

		if err != nil {
			fmt.Println(err)
		}

		if orderStatus.Status != "" {
			intStatus := blackboxservice.FromStringStatusToInt(orderStatus.Status)
			orderService.orderRepo.UpdateOrderStatus(order.Id, intStatus)
		}
	}
}
