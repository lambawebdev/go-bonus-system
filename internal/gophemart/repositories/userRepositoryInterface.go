package repositories

import (
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

type UserRepo interface {
	GetExistingUser(login string) (entities.User, error)
	CheckIfUserLoginAlreadyExists(login string) (bool, error)
	CreateUser(userDto dto.User) (entities.User, error)
}
