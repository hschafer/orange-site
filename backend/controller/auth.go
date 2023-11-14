package controller

import (
	"net/http"
	"orange-site/backend/model"

	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	login := new(LoginInfo)
	err := c.Bind(login)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	authenticated, err := model.AuthenticateUser(login.Username, login.Password)

	return c.JSON(http.StatusOK, authenticated)
}
