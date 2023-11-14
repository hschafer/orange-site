package controller

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidToken(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)

	if ok {
		return errors.New("Authorization token missing or invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Failed to convert claims")
	}

	return c.JSON(http.StatusOK, claims)

}
