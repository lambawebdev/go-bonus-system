package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/validators"
)

type OrderHandler struct {
	orderRepo repositories.OrderRepo
}

func NewOrderHandler(orderRepo repositories.OrderRepo) *OrderHandler {
	return &OrderHandler{
		orderRepo: orderRepo,
	}
}

func (orderHandler *OrderHandler) GetOrders(res http.ResponseWriter, req *http.Request) {
	orders, err := orderHandler.orderRepo.GetOrders(req.Context())

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(orders)
}

func (orderHandler *OrderHandler) CreateOrder(res http.ResponseWriter, req *http.Request) {
	scanner := bufio.NewScanner(req.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}

	defer buf.Reset()

	err := validators.ValidateOrderRequest(buf.String())

	if err != nil {
		http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	exists, err := orderHandler.orderRepo.CheckIfOrderWasAddedByAnotherUser(req.Context(), buf.String())

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(res, "number was used by another user", http.StatusConflict)
		return
	}

	_, err = orderHandler.orderRepo.CreateOrder(req.Context(), buf.String())
	statusOk := http.StatusAccepted

	if err != nil {
		statusOk, err = handlePGSqlError(err)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("content-Type", "application/json")
	res.WriteHeader(statusOk)
}

func handlePGSqlError(err error) (int, error) {
	var pgerr *pgconn.PgError

	if errors.As(err, &pgerr) {
		if pgerr.Code == pgerrcode.UniqueViolation {
			return http.StatusOK, nil
		}
	}

	return 500, err
}
