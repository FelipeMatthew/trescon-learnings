package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEvents(c echo.Context) error {
	return c.String(http.StatusOK, "get events")
}
