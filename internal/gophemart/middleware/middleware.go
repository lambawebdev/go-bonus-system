package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lambawebdev/go-bonus-system/internal/gophemart/services/jwtservice"
)

type CtxKey string
var UserIDkey CtxKey = "user_id"

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearer := req.Header.Get("Authorization")

		if bearer == "" || !strings.HasPrefix(bearer, "Bearer ") {
			http.Error(res, "Missing authentication token", http.StatusUnauthorized)
			return
		}

		jwt := strings.Split(bearer, "Bearer ")[1]

		userID := jwtservice.GetUserID(jwt)

		if userID == -1 {
			http.Error(res, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), &UserIDkey, userID)
		newReq := req.WithContext(ctx)

		h.ServeHTTP(res, newReq)
	})
}
