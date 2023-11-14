package controller

import (
	"net/http"
	"orange-site/backend/model"

	"github.com/labstack/echo/v4"
)

func GetComments(c echo.Context) error {
	id := c.Param("id")
	post, err := model.GetComments(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	} else {
		return c.JSON(http.StatusOK, post)
	}
}
