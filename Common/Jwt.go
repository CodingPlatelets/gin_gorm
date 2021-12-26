package Common

import (
	"strconv"
	"time"

	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("shadfjkahsdkjfhn2qjsahfjk&*&*24n1kjk")

func ReleaseToken(user Model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ID:        strconv.Itoa(int(user.UserId)),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "Todo-backend",
		Subject:   "user token",
		Audience:  []string{"Todo"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
