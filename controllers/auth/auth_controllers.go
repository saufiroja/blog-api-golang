package auth

import (
	"echo/blog-api/entity"
	"echo/blog-api/service/auth"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S auth.AuthService
}

/*
	Register a new user
	@Summary Register a new user
	@Router POST /api/register
	@Success 200 {object} entity.User
	@Failure 400 {object} entity.Error
	@Failure 500 {object} entity.Error
*/
func (c *Controller) Register(ctx echo.Context) error {
	// model user
	user := entity.User{
		ID: uuid.New().String(),
	}
	// bind user
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// register user
	users, er := c.S.Register(user)

	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": "success register",
		"data":    users,
	})
}

/*
	Login a user
	@Summary Login a user
	@Router POST /api/login
	@Success 200 {object} entity.User
	@Failure 400 {object} entity.Error
	@Failure 500 {object} entity.Error
*/
func (c *Controller) Login(ctx echo.Context) error {
	// model user
	user := entity.User{}
	// bind user
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// login user
	token, er := c.S.Login(user.Email, user.Password)

	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": er.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success login",
		"data":    token,
	})
}
