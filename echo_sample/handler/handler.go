package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello world"+username)
	}
}
