package controller

import (
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	e.GET("/posts", GetAllPosts)
	e.GET("/post/:id", GetPost)
}
