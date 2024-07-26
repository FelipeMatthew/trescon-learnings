package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)



func GetUserById(c echo.Context) error {
	userId := c.Param("userId")

	for _, user := range models.Users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.String(http.StatusBadRequest, "User not found")
}