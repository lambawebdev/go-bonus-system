package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

var balance = entities.Balance{Current: 100.40, Withdrawn: 50.25}
var byteBalance, _ = json.Marshal(balance)

var withdrawal = entities.Withdrawal{
	Order:       OrderNumber,
	Sum:         125.25,
	ProcessedAt: time.Now().Local(),
}
var withdrawals = []entities.Withdrawal{withdrawal}
var byteWithdrawals, _ = json.Marshal(withdrawals)

var transactionDto = dto.Transaction{Number: OrderNumber, Amount: 30}
var byteTransactionDto, _ = json.Marshal(transactionDto)

func TestGetBalance(t *testing.T) {
	t.Run("GetBalance", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/user/balance", nil)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		transRepoMock := mocks.NewMockTransRepo(mockCtrl)
		orderRepoMock := mocks.NewMockOrderRepo(mockCtrl)
		withdrawalRepoMock := mocks.NewMockWithdrawalRepo(mockCtrl)

		balanceHandler := NewBalanceHandler(orderRepoMock, transRepoMock, withdrawalRepoMock)

		transRepoMock.EXPECT().GetBalance(r.Context()).Return(balance, nil)

		balanceHandler.GetBalance(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.JSONEq(t, string(byteBalance), w.Body.String())
	})
}

func TestGetWithdrawals(t *testing.T) {
	t.Run("GetWithdrawals", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/user/withdrawals", nil)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		transRepoMock := mocks.NewMockTransRepo(mockCtrl)
		orderRepoMock := mocks.NewMockOrderRepo(mockCtrl)
		withdrawalRepoMock := mocks.NewMockWithdrawalRepo(mockCtrl)

		balanceHandler := NewBalanceHandler(orderRepoMock, transRepoMock, withdrawalRepoMock)

		withdrawalRepoMock.EXPECT().GetWithdrawals(r.Context()).Return(withdrawals, nil)

		balanceHandler.GetWithdrawals(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.JSONEq(t, string(byteWithdrawals), w.Body.String())
	})
}

func TestCreateWithdrawal(t *testing.T) {
	t.Run("CreateWithdrawal", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/balance/withdraw", strings.NewReader(string([]byte(byteTransactionDto))))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		transRepoMock := mocks.NewMockTransRepo(mockCtrl)
		orderRepoMock := mocks.NewMockOrderRepo(mockCtrl)
		withdrawalRepoMock := mocks.NewMockWithdrawalRepo(mockCtrl)

		balanceHandler := NewBalanceHandler(orderRepoMock, transRepoMock, withdrawalRepoMock)

		withdrawalRepoMock.EXPECT().CreateWithdrawal(r.Context(), transactionDto).Return(nil)

		balanceHandler.Withdraw(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	})
}
