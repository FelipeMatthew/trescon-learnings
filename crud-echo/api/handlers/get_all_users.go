package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Users)
}