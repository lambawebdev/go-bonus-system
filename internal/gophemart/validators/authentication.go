package validators

import (
	"errors"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/bcryptService"
)

func ValidateAuthRequest(userRepo repositories.UserRepo, userDto *dto.User) error {
	if userDto.Login == "" {
		return errors.New("login must be present")
	}

	if userDto.Password == "" {
		return errors.New("password must be present")
	}

	user, err := userRepo.GetExistingUser(userDto.Login)

	if err != nil {
		return err
	}

	err = bcryptService.ValidateUserPassword(user.Password, userDto.Password)

	if err != nil {
		return err
	}

	return nil
}
