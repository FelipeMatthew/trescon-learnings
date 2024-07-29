package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFiles(c echo.Context) error {
	return c.String(http.StatusOK, "hi from get files")
}
