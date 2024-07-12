package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)


func DeleteUserById(c echo.Context) error {
	userId := c.Param("userId")
	users := models.Users

	index := -1

	for idx, user := range users {
		if user.Id == userId {
			index = idx
			break
		}
	}

	// User not founded
	if index == -1 {
		return c.String(http.StatusBadRequest, "User not founded")
	}

	// User founded
	users = append(users[:index], users[index +1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "User deleted",
		"users": users,
	})
}