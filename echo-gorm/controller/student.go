package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handlers
func GetStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Get Students")
}