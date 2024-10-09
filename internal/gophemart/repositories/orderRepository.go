package repositories

import (
	"context"
	"database/sql"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

const GET_ORDERS = `
    SELECT id, user_id, number, status, created_at, updated_at FROM orders WHERE user_id = $1
	ORDER BY created_at DESC
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
