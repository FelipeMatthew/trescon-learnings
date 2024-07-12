package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}

func HandlerUser(c echo.Context) error {
	userId := c.Param("userId")
	return c.String(http.StatusOK, userId)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Users)
}