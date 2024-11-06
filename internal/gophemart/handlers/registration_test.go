package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/dto"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories/mocks"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/bcryptservice"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/jwtservice"
	"github.com/stretchr/testify/assert"
)

const (
	Password = "1234"
	Login    = "john.doe"
)

var hashPassword, _ = bcryptservice.HashPassword(Password)

var userDto = dto.User{Login: "john.doe", Password: Password}
var user = entities.User{ID: 1, Login: userDto.Login, Password: hashPassword}

var userDtoBytes, _ = json.Marshal(userDto)

func TestRegister(t *testing.T) {
	t.Run("NewUser", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/register", strings.NewReader(string(userDtoBytes)))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockUserRepo(mockCtrl)
		regHandler := NewRegistrationHandler(m)

		m.EXPECT().CheckIfUserLoginAlreadyExists("john.doe").Return(false, nil)
		m.EXPECT().CreateUser(userDto).Return(user, nil)

		regHandler.Register(w, r)

		bearer := w.Header().Get("Authorization")
		jwt := strings.Split(bearer, "Bearer ")[1]

		userID := jwtservice.GetUserID(jwt)

		assert.Equal(t, user.ID, userID)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	})
}

func TestAuth(t *testing.T) {
	t.Run("AuthSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(string(userDtoBytes)))

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		m := mocks.NewMockUserRepo(mockCtrl)
		authHandler := NewAuthenticationHandler(m)

		m.EXPECT().GetExistingUser(userDto.Login).Return(user, nil).Times(2)

		authHandler.Authenticate(w, r)

		bearer := w.Header().Get("Authorization")
		jwt := strings.Split(bearer, "Bearer ")[1]

		userID := jwtservice.GetUserID(jwt)

		assert.Equal(t, user.ID, userID)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	})
}
