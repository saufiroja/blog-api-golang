package auth

import (
	"echo/blog-api/config"
	c "echo/blog-api/controllers/auth"
	r "echo/blog-api/repository/auth"
	s "echo/blog-api/service/auth"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, conf config.Config) {
	// init db
	db := config.InitDB(conf)

	// init repository
	repo := r.NewAuthRepository(db)
	service := s.NewAuthService(repo, conf)
	control := c.Controller{
		S: service,
	}

	// group
	g := e.Group("/api")

	// register user
	g.POST("/register", control.Register)
	g.POST("/login", control.Login)
}
