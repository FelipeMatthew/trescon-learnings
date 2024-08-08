package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/FelipeMatthew/go-learnings/crud-echo/api/utils"
	"github.com/labstack/echo/v4"
)

func UpdateByUserId(c echo.Context) error {
	userId := c.Param("userId")

	userFound, err := utils.FindUserById(userId)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	// Bind
	userUpdated := models.UserType{}
	err = c.Bind(&userUpdated)
	if err != nil {
		return c.String(http.StatusBadRequest, "Body user not found")
	}

	// Update
	userFound.Name = userUpdated.Name
	userFound.IsActive = userUpdated.IsActive
	userUpdated.Id = userId

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "user updated",
		"user":   userUpdated,
	})

}
