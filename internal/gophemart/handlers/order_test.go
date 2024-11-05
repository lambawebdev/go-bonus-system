package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	OrderNumber = "354056509"
	WrongNumber = "111111111"
)

var order = entities.Order{
	ID:        1,
	UserID:    1,
	Number:    OrderNumber,
	Status:    "NEW",
	CreatedAt: time.Now().Local(),
	UpdatedAd: time.Now().Local(),
	Accrual:   0,
}

var orders = []entities.Order{order}
var byteOrders, _ = json.Marshal(orders)

func TestCreateOrder(t *testing.T) {
	t.Run("CreateOrder", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(string([]byte(OrderNumber))))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockOrderRepo(mockCtrl)
		orderHandler := NewOrderHandler(m)

		m.EXPECT().CheckIfOrderWasAddedByAnotherUser(r.Context(), OrderNumber).Return(false, nil)
		m.EXPECT().CreateOrder(r.Context(), OrderNumber).Return(order, nil)

		orderHandler.CreateOrder(w, r)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	})

	t.Run("CreateOrderWhenNumberIsUsed", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(string([]byte(OrderNumber))))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockOrderRepo(mockCtrl)
		orderHandler := NewOrderHandler(m)

		m.EXPECT().CheckIfOrderWasAddedByAnotherUser(r.Context(), OrderNumber).Return(true, nil)

		orderHandler.CreateOrder(w, r)

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "number was used by another user\n", w.Body.String())
	})

	t.Run("CreateOrderWhenNumberIsWrong", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/orders", strings.NewReader(string([]byte(WrongNumber))))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockOrderRepo(mockCtrl)
		orderHandler := NewOrderHandler(m)

		orderHandler.CreateOrder(w, r)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
		assert.Equal(t, "wrong number\n", w.Body.String())
	})
}

func TestGetOrders(t *testing.T) {
	t.Run("GetOrders", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/user/orders", nil)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockOrderRepo(mockCtrl)
		orderHandler := NewOrderHandler(m)

		m.EXPECT().GetOrders(r.Context()).Return(orders, nil)

		orderHandler.GetOrders(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.JSONEq(t, string(byteOrders), w.Body.String())
	})
}
