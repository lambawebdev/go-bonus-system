package repositories

import (
	"context"
	"database/sql"
	"math"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

const GET_ORDERS = `
    SELECT orders.id, orders.user_id, orders.number, orders.status, orders.created_at, orders.updated_at, COALESCE(transactions.amount, 0)
	FROM orders
	LEFT JOIN transactions ON orders.id = transactions.order_id
	WHERE orders.user_id = $1
	ORDER BY orders.created_at DESC
    `

const GET_NOT_PROCESSED_ORDERS = `
    SELECT id, user_id, number, status, created_at, updated_at FROM orders WHERE status != 3
	ORDER BY created_at DESC
    `

const UPDATE_ORDER_STATUS = `
	UPDATE orders SET status = $1 WHERE id = $2
`

const GET_ORDER = `
    SELECT id, user_id, number, status, created_at, updated_at FROM orders WHERE number = $1
    `

const ORDER_WAS_ADDED_BY_ANOTHER_USER = `
    SELECT EXISTS(SELECT * FROM orders WHERE number = $1 AND user_id != $2)
    `

const CREATE_ORDER = `
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
	rows, err := repository.database.Query(GET_ORDERS, ctx.Value("user_id"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var amount float64
		var status int
		var order entities.Order

		if err := rows.Scan(&order.Id, &order.UserId, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd, &amount); err != nil {
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
	rows, err := repository.database.Query(GET_NOT_PROCESSED_ORDERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var status int
		var order entities.Order

		if err := rows.Scan(&order.Id, &order.UserId, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
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

func (repository *OrderRepository) UpdateOrderStatus(orderId int, status int) error {
	_, err := repository.database.Exec(UPDATE_ORDER_STATUS, status, orderId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetOrderByNumber(ctx context.Context, number string) (entities.Order, error) {
	var status int
	var order entities.Order

	if err := repository.database.QueryRow(GET_ORDER, number).
		Scan(&order.Id, &order.UserId, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
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

	if err := repository.database.QueryRow(ORDER_WAS_ADDED_BY_ANOTHER_USER, number, ctx.Value("user_id")).Scan(&exists); err != nil {
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

	if err := repository.database.QueryRow(CREATE_ORDER, number, ctx.Value("user_id")).
		Scan(&order.Id, &order.UserId, &order.Number, &status, &order.CreatedAt, &order.UpdatedAd); err != nil {
		if err == sql.ErrNoRows {
			return order, err
		}

		return order, err
	}

	order.Status = entities.TransformStatusToString(status)

	return order, nil
}
