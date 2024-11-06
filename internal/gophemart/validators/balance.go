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

	if !luhnAlgorithm(transactionDto.Number) {
		return errors.New("wrong number")
	}

	return nil
}

func LuhnAlgorithm(cardNumber string) bool {
	total := 0
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit

		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0
}
