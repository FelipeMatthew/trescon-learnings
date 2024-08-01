package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdmins(c echo.Context) error {
	return c.String(http.StatusOK, "Hello im working on admins")
}
