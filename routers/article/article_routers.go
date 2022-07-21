package article

import (
	"echo/blog-api/config"
	c "echo/blog-api/controllers/article"
	r "echo/blog-api/repository/article"
	s "echo/blog-api/service/article"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ArticleRoutes(echo *echo.Echo, conf config.Config) {
	db := config.InitDB(conf)

	repo := r.NewArticleRepository(db)
	service := s.NewArticleService(repo, conf)
	control := c.Controller{
		S: service,
	}

	e := echo.Group("/api")
	g := echo.Group("/api")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(conf.JWT_SECRET),
		SigningMethod: "HS256",
	}))

	g.POST("/articles", control.CreateArticle)
	e.GET("/articles", control.FindAllArticle)
	e.GET("/articles/:id", control.FindByIDArticle)
	g.PUT("/articles/:id", control.UpdateArticle)
	g.DELETE("/articles/:id", control.DeleteArticle)
}
