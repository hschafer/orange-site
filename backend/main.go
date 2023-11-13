package main

import (
	"fmt"
	"log"
	"net/http"

	"orange-site/backend/model"
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
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("HI THERE I GOT a post")
	fmt.Print(posts)

	return c.JSON(http.StatusOK, posts)
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

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
