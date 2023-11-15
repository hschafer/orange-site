package main

import (
	"orange-site/backend/controller"
	"orange-site/backend/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"

	_ "github.com/lib/pq"
)

// Setup
func main() {
	storage.InitDBConnection()

	// Echo instance
	e := echo.New()

	// Get server secret to set up JWT middleware
	secret := controller.GetSecret()
	jwt := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		Skipper:    controller.SkipRoutes,
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: []byte(secret),
	//	Skipper:    controller.SkipRoutes,
	//}))

	// Routes
	controller.SetRoutes(e, jwt)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
