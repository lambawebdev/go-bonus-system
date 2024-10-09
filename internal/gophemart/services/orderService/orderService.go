package orderservice

import (
	"fmt"
	"time"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	blackboxservice "github.com/lambawebdev/go-bonus-system/internal/gophemart/services/blackBoxService"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
	transRepo repositories.TransRepo
}

func NewOrderService(orderRepo repositories.OrderRepo, transRepo repositories.TransRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		transRepo: transRepo,
	}
}

func (orderServices *OrderService) RunUpdateOrdersStatuses() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			orderServices.UpdateOrdersStatuses()
		}
	}
}

func (orderService *OrderService) UpdateOrdersStatuses() {
	orders, err := orderService.orderRepo.GetNotProcessedOrders()

	if err != nil {
		fmt.Println(err)
	}

	for _, order := range orders {
		orderAccrual, err := blackboxservice.GetOrderAccrual(order.Number)

		if err != nil {
			fmt.Println(err)
		}

		orderAccrual.OrderId = order.Id

		if orderAccrual.Status != "" {
			intStatus := blackboxservice.FromStringStatusToInt(orderAccrual.Status)
			orderService.orderRepo.UpdateOrderStatus(order.Id, intStatus)

			if intStatus == entities.STATUS_PROCESSED {
				orderService.transRepo.CreateTransaction(order.UserId, orderAccrual)
			}
		}
	}
}
