package main

import (
	"echo/blog-api/config"
	"echo/blog-api/routers/auth"
	"echo/blog-api/routers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	conf := config.Config{}

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// ROUTES
	auth.AuthRoutes(e, conf)
	user.UserRoutes(e, conf)

	e.Logger.Fatal(e.Start("127.0.0.1:4000"))
}
