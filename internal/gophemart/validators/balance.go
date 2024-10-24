package validators

import (
	"errors"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
)

func ValidateBalanceRequest(transactionDto *dto.Transaction) error {
	if transactionDto.Number == "" {
		return errors.New("number must be present")
	}

	if transactionDto.Amount == 0 {
		return errors.New("amount must be present")
	}

	return nil
}
