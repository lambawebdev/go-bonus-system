package repositories

import (
	"context"
	"database/sql"
	"math"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/middleware"
)

const GetOrders = `
    SELECT orders.id, orders.user_id, orders.number, orders.status, orders.created_at, orders.updated_at, COALESCE(transactions.amount, 0)
	FROM orders
	LEFT JOIN transactions ON orders.id = transactions.order_id
	WHERE orders.user_id = $1
	ORDER BY orders.created_at DESC
    `

const GetNotProcessedOrders = `
    SELECT id, user_id, number, status, created_at, updated_at FROM orders WHERE status != 3
	ORDER BY created_at DESC
    `

const UpdateOrderStatus = `
	UPDATE orders SET status = $1 WHERE id = $2
`

const GetOrder = `
    SELECT id, user_id, number, status, created_at, updated_at FROM orders WHERE number = $1
    `

const OrderWasAddedByAnotherUser = `
    SELECT EXISTS(SELECT * FROM orders WHERE number = $1 AND user_id != $2)
    `

const CreateOrder = `
    INSERT INTO orders (number, user_id) VALUES ($1, $2) 
	RETURNING id, user_id, number, status, created_at, updated_at
	`

type OrderRepository struct {
	database *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		database: db,
	}
}

func (repository *OrderRepository) GetOrders(ctx context.Context) ([]entities.Order, error) {
	rows, err := repository.database.Query(GetOrders, ctx.Value(&middleware.UserIDkey))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var amount float64
		var status int
		var order entities.Order

		if err := rows.Scan(&order.ID, &order.UserID, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd, &amount); err != nil {
			return orders, err
		}

		if amount > 0 {
			amount = math.Abs(amount) / 10000
		}

		order.Accrual = amount
		order.Status = entities.TransformStatusToString(status)
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (repository *OrderRepository) GetNotProcessedOrders() ([]entities.Order, error) {
	rows, err := repository.database.Query(GetNotProcessedOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var status int
		var order entities.Order

		if err := rows.Scan(&order.ID, &order.UserID, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
			return orders, err
		}

		order.Status = entities.TransformStatusToString(status)
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}

func (repository *OrderRepository) UpdateOrderStatus(orderID int, status int) error {
	_, err := repository.database.Exec(UpdateOrderStatus, status, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetOrderByNumber(ctx context.Context, number string) (entities.Order, error) {
	var status int
	var order entities.Order

	if err := repository.database.QueryRow(GetOrder, number).
		Scan(&order.ID, &order.UserID, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
		if err == sql.ErrNoRows {
			return order, err
		}

		return order, err
	}

	order.Status = entities.TransformStatusToString(status)

	return order, nil
}

func (repository *OrderRepository) CheckIfOrderWasAddedByAnotherUser(ctx context.Context, number string) (bool, error) {
	var exists bool

	if err := repository.database.QueryRow(OrderWasAddedByAnotherUser, number, ctx.Value(&middleware.UserIDkey)).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return exists, nil
		}

		return exists, err
	}

	return exists, nil
}

func (repository *OrderRepository) CreateOrder(ctx context.Context, number string) (entities.Order, error) {
	var status int
	var order entities.Order

	if err := repository.database.QueryRow(CreateOrder, number, ctx.Value(&middleware.UserIDkey)).
		Scan(&order.ID, &order.UserID, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
		if err == sql.ErrNoRows {
			return order, err
		}

		return order, err
	}

	order.Status = entities.TransformStatusToString(status)

	return order, nil
}
