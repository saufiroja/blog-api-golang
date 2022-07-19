package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	ID uint
	// StandardClaims contains the standard jwt claims.
	jwt.StandardClaims
}

func CreateToken(id uint, secret string) (string, error) {
	claims := JWTClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
