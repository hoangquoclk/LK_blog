package service

import (
	"LK_blog/model"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(role string, email string) (string, error, time.Time) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &model.Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authSecret := []byte(os.Getenv("AUTH_SECRET"))
	tokenString, err := token.SignedString(authSecret)

	return tokenString, err, expirationTime
}
