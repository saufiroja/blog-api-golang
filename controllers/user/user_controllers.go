package user

import (
	"echo/blog-api/service/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S user.UserService
}

func (c *Controller) FindAllUsers(ctx echo.Context) error {
	users, err := c.S.FindAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success find all users",
		"data":    users,
	})
}

func (c *Controller) FindByIDUsers(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := c.S.FindByIDUsers(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success find by id users",
		"data":    user,
	})
}
