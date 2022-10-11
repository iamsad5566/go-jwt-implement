package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func GenerateToken(usr User) (string, error) {
	maxAge := 60 * 60 * 24
	expiresAt := time.Now().Add(time.Duration(maxAge) * time.Second)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   usr.Account,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		UserID: 1,
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (int, string, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRETKEY")), nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	id := claims.UserID
	userName := claims.Subject
	return id, userName, nil
}
