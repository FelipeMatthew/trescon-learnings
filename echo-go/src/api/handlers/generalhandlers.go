package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GettingStart(c echo.Context) error {
	return c.String(http.StatusOK, "My first echo go lang system :)")
}