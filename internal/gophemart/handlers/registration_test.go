package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()
	t.Run("NewUser", func(t *testing.T) {
		loginNotExistsRow := sqlmock.NewRows([]string{"login"}).
			AddRow(false)

		createUserRow := sqlmock.NewRows([]string{"id", "login"}).
			AddRow("0", "john.doe")

		mock.ExpectQuery(repositories.CHECK_LOGIN_FOR_EXISTANCE).WithArgs("john.doe").WillReturnRows(loginNotExistsRow)
		mock.ExpectQuery(repositories.CREATE_USER).WillReturnRows(createUserRow)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/user/register", strings.NewReader(`{"login": "john.doe", "password": "1234"}`))

		ctx := context.WithValue(r.Context(), "DB", db)
		r = r.WithContext(ctx)

		userRepo := repositories.NewUserRepository(db)
		regHandler := NewRegistrationHandler(userRepo)

		regHandler.Register(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("not all expectations were met: %v", err)
		}
	})
}
