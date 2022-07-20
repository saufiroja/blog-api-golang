package user

import (
	"echo/blog-api/entity"
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

func (c *Controller) UpdateUser(ctx echo.Context) error {
	// entity user
	user := &entity.User{}
	// param id
	id, _ := strconv.Atoi(ctx.Param("id"))
	ID := uint(id)
	// req body
	err := ctx.Bind(user)
	// check if req body invalid
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// update user
	er := c.S.UpdateUser(ID, *user)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}
	// return success update user
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success update user",
	})
}

func (c *Controller) DeleteUser(ctx echo.Context) error {
	// entity user
	user := &entity.User{}
	// param id
	id, _ := strconv.Atoi(ctx.Param("id"))

	// delete user
	err := c.S.DeleteUser(uint(id), *user)
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
