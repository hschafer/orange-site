package controller

import (
	"github.com/labstack/echo/v4"
)

func SkipRoutes(c echo.Context) bool {
	noToken := map[string]bool{
		"/posts":             true,
		"/post/:id":          true,
		"/post/:id/comments": true,
		"/login":             true,
	}
	return noToken[c.Request().URL.Path]
}

func SetRoutes(e *echo.Echo, jwt echo.MiddlewareFunc) {
	e.GET("/posts", GetAllPosts)
	e.GET("/post/:id", GetPost)
	e.GET("/post/:id/comments", GetComments)
	e.POST("/post/new", NewPost, jwt)
	e.POST("/login", Login)
	e.POST("/register", Register)
	e.POST("/comment/new", PostComment, jwt)
}
