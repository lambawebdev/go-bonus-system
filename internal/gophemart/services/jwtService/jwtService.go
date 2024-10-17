package jwtService

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/config"
)

const TOKEN_EXP = time.Hour * 3

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

func BuildJWTString(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		UserID: userId,
	})

	tokenString, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserId(tokenString string) int {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecret()), nil
		})
	if err != nil {
		return -1
	}

	if !token.Valid {
		return -1
	}

	return claims.UserID
}

func CreateCookieWithJWT(jwt string) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		MaxAge:   36000,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
