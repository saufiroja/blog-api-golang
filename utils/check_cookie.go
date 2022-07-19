package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"error":   err.Error(),
					"message": "cookie not found",
				})
			}
			log.Println(err)
			return err
		}

		if cookie.Value == "" {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		}

		return next(c)
	}
}
