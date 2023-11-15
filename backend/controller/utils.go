package controller

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetSecret() string {
	secret := os.Getenv("SERVER_SECRET")
	if secret == "" {
		panic("No secret specified")
	}
	return secret
}

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix() // One hour timeout

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := GetSecret()
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func ParseToken(c echo.Context) (jwt.MapClaims, error) {
	// from https://github.com/labstack/echo-jwt
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}
	return claims, nil
}
