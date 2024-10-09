package validators

import (
	"errors"
	"fmt"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
)

func ValidateRegisterRequest(userRepo repositories.UserRepo, userDto *dto.User) error {
	if userDto.Login == "" {
		return errors.New("login must be present") //go validate
	}

	if userDto.Password == "" {
		return errors.New("password must be present")
	}

	loginExists, err := userRepo.CheckIfUserLoginAlreadyExists(userDto.Login)

	if err != nil {
		fmt.Println(err)

		return err
	}

	if loginExists {
		return errors.New("login already exists")
	}

	return nil
}
