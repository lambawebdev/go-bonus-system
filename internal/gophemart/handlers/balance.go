package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/validators"
)

type BalanceHandler struct {
	orderRepo      repositories.OrderRepo
	transRepo      repositories.TransRepo
	withdrawalRepo repositories.WithdrawalRepo
}

func NewBalanceHandler(orderRepo repositories.OrderRepo, transRepo repositories.TransRepo, withdrawalRepo repositories.WithdrawalRepo) *BalanceHandler {
	return &BalanceHandler{
		orderRepo:      orderRepo,
		transRepo:      transRepo,
		withdrawalRepo: withdrawalRepo,
	}
}

func (balanceHandler *BalanceHandler) GetBalance(res http.ResponseWriter, req *http.Request) {
	balance, err := balanceHandler.transRepo.GetBalance(req.Context())

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(balance)
}

func (balanceHandler *BalanceHandler) GetWithdrawals(res http.ResponseWriter, req *http.Request) {
	withdrawals, err := balanceHandler.withdrawalRepo.GetWithdrawals(req.Context())

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(withdrawals)
}

func (balanceHandler *BalanceHandler) Withdraw(res http.ResponseWriter, req *http.Request) {
	var transactionDto dto.Transaction

	err := json.NewDecoder(req.Body).Decode(&transactionDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validators.ValidateBalanceRequest(&transactionDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = balanceHandler.withdrawalRepo.CreateWithdrawal(req.Context(), transactionDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
