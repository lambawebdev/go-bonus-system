package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/jwtservice"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/validators"
)

type AuthenticationHandler struct {
	userRepo repositories.UserRepo
}

func NewAuthenticationHandler(userRepo repositories.UserRepo) *AuthenticationHandler {
	return &AuthenticationHandler{
		userRepo: userRepo,
	}
}

func (authHandler *AuthenticationHandler) Authenticate(res http.ResponseWriter, req *http.Request) {
	var userDto dto.User

	err := json.NewDecoder(req.Body).Decode(&userDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validators.ValidateAuthRequest(authHandler.userRepo, &userDto)

	if err != nil {
		http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := authHandler.userRepo.GetExistingUser(userDto.Login)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := jwtservice.BuildJWTString(user.ID)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Authorization", "Bearer "+jwt)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
