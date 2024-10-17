package validators

import (
	"errors"
)

func ValidateOrderRequest(number string) error {
	if number == "" {
		return errors.New("number must be present")
	}

	if !luhnAlgorithm(number) {
		return errors.New("wrong number")
	}

	return nil
}

func luhnAlgorithm(cardNumber string) bool {
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
