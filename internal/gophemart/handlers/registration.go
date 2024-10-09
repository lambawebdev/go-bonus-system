package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/jwtService"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/validators"
)

type RegistrationHandler struct {
	userRepo repositories.UserRepo //interface
}

func NewRegistrationHandler(userRepo repositories.UserRepo) *RegistrationHandler {
	return &RegistrationHandler{
		userRepo: userRepo,
	}
}

func (regHandler *RegistrationHandler) Register(res http.ResponseWriter, req *http.Request) {
	var userDto dto.User

	err := json.NewDecoder(req.Body).Decode(&userDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validators.ValidateRegisterRequest(regHandler.userRepo, &userDto)

	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := regHandler.userRepo.CreateUser(userDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := jwtService.BuildJWTString(user.Id)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Authorization", "Bearer "+jwt)
	res.Header().Set("content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
