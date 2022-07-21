package article

import (
	"echo/blog-api/entity"
	"echo/blog-api/service/article"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S article.ArticleService
}

/*
	CreateArticle
	@Summary Create an article
	@Router POST /api/articles
	@Success 201 {object} entity.Article
	@Failure 400 {object} entity.Error
	@Failure 500 {object} entity.Error
	@Security JWT
*/
func (c *Controller) CreateArticle(ctx echo.Context) error {
	// check current user
	user := ctx.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	// model article
	article := entity.Article{
		ID:     uuid.New().String(),
		UserID: userID,
	}
	// bind article
	err := ctx.Bind(&article)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// create article
	articles, er := c.S.CreateArticle(article)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": "success create article",
		"data":    articles,
	})

}

/*
	FindAllArticle
	@Summary Find all article
	@Router GET /api/articles
	@Success 200 {object} entity.Article
	@Failure 500 {object} entity.Error
*/
func (c *Controller) FindAllArticle(ctx echo.Context) error {
	// find all article
	articles, er := c.S.FindAllArticle()
	// if articles is empty, return error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}
	// return articles
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success find all article",
		"data":    articles,
	})
}

/*
	FindByIDArticle
	@Summary Find an article by id
	@Router GET /api/articles/:id
	@Success 200 {object} entity.Article
	@Failure 500 {object} entity.Error
*/
func (c *Controller) FindByIDArticle(ctx echo.Context) error {
	// get id from url
	id := ctx.Param("id")
	// find article by id
	articles, er := c.S.FindByIDArticle(id)
	// if article is empty, return error
	if er != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": er.Error(),
		})
	}
	// return article
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success find article by id",
		"data":    articles,
	})
}

/*
	UpdateArticle
	@Summary Update an article
	@RequestBody entity.Article
	@Router PUT /api/articles/:id
	@Success 200 {object} entity.Article
	@Failure 400 {object} entity.Error
	@Failure 500 {object} entity.Error
	@Security JWT
*/
func (c *Controller) UpdateArticle(ctx echo.Context) error {
	// model article
	article := entity.Article{}

	// get id from url
	id := ctx.Param("id")

	// bind article
	err := ctx.Bind(&article)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// update article
	articles, er := c.S.UpdateArticle(id, article)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success update article",
		"data":    articles,
	})
}

/*
	DeleteArticle
	@Summary Delete an article
	@Router DELETE /api/articles/:id
	@Success 200 {object} entity.Article
	@Failure 500 {object} entity.Error
	@Security JWT
*/
func (c *Controller) DeleteArticle(ctx echo.Context) error {
	// entity article
	article := entity.Article{}

	// param id
	id := ctx.Param("id")
	// delete article
	er := c.S.DeleteArticle(id, article)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}
	// return article
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success delete article",
	})
}
