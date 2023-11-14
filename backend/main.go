package main

import (
	"os"

	"orange-site/backend/controller"
	"orange-site/backend/storage"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func getSecret() string {
	secret := os.Getenv("SERVER_SECRET")
	if secret == "" {
		panic("No secret specified")
	}
	return secret
}

// Setup
func main() {
	storage.InitDBConnection()

	// Echo instance
	e := echo.New()

	// Get server secret to set up JWT middleware
	secret := getSecret()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		Skipper:    controller.SkipRoutes,
	}))

	// Routes
	controller.SetRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
