package user

import (
	"echo/blog-api/config"
	c "echo/blog-api/controllers/user"
	r "echo/blog-api/repository/user"
	s "echo/blog-api/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(echo *echo.Echo, conf config.Config) {
	// init DB
	db := config.InitDB(conf)

	// init repository
	repo := r.NewUserRepository(db)
	service := s.NewUserService(repo, conf)
	control := c.Controller{
		S: service,
	}

	// group
	g := echo.Group("/api")

	// middlewares
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(conf.JWT_SECRET),
		SigningMethod: "HS256",
	}))

	g.GET("/users", control.FindAllUsers)
	g.GET("/users/:id", control.FindByIDUsers)
}
