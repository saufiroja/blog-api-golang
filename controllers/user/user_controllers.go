package user

import (
	"echo/blog-api/entity"
	"echo/blog-api/service/user"
	"net/http"

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
	id := ctx.Param("id")

	user, err := c.S.FindByIDUsers(id)
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

func (c *Controller) UpdateUser(ctx echo.Context) error {
	// entity user
	user := entity.User{}
	// param id
	id := ctx.Param("id")
	// request body
	er := ctx.Bind(&user)
	if er != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": er.Error(),
		})
	}

	// update user
	err := c.S.UpdateUser(id, user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed update user",
			"error":   err.Error(),
		})
	}
	// return success update user
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success update user",
	})
}

func (c *Controller) DeleteUser(ctx echo.Context) error {
	// entity user
	user := entity.User{}
	// param id
	id := ctx.Param("id")

	// delete user
	err := c.S.DeleteUser(id, user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}
	// return success delete user
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success delete user",
	})
}
