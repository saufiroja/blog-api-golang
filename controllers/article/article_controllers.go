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

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success create article",
		"data":    articles,
	})

}
