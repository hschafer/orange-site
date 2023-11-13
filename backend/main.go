package main

import (
	"fmt"
	"net/http"

	"orange-site/backend/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func posts(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World2!")
}

// Setup
func main() {
	storage.InitDBConnection()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/posts", posts)

	//// Test DB
	//people := []User{}
	//db.Select(&people, "SELECT * FROM users;")
	//fmt.Print(people)

	posts := []storage.Post{}
	storage.GetDBConnection().Select(&posts, "SELECT * FROM posts;")
	fmt.Print(posts)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
