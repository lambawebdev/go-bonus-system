package repositories

import (
	"database/sql"
	"errors"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/bcryptservice"
)

const CreateUser = `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id, login`
const GetUser = `SELECT id, login, password FROM users WHERE login = $1`
const CheckLoginForExistance = `SELECT EXISTS(SELECT * FROM users WHERE login = $1)`

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (repository *UserRepository) GetExistingUser(login string) (entities.User, error) {
	var user entities.User

	if err := repository.database.QueryRow(GetUser, login).Scan(&user.ID, &user.Login, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not exists")
		}

		return user, err
	}

	return user, nil
}

func (repository *UserRepository) CheckIfUserLoginAlreadyExists(login string) (bool, error) {
	var exists bool

	if err := repository.database.QueryRow(CheckLoginForExistance, login).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return exists, nil
		}

		return exists, err
	}

	return exists, nil
}

func (repository *UserRepository) CreateUser(userDto dto.User) (entities.User, error) {
	hashedPass, _ := bcryptservice.HashPassword(userDto.Password)
	userDto.Password = hashedPass

	var user entities.User

	if err := repository.database.QueryRow(CreateUser, userDto.Login, userDto.Password).Scan(&user.ID, &user.Login); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}

		return user, err
	}

	return user, nil
}
