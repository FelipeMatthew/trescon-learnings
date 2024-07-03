package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "My first middleware using go")
}