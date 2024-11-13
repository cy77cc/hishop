package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Openid string `json:"openid"`
	jwt.RegisteredClaims
}

func ParseJWT(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if cliam, ok := token.Claims.(*Claims); ok && token.Valid {
		return cliam, err
	} else {
		return nil, err
	}
}
