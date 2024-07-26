package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	users := models.Users
	newUSer := models.UserType{}

	// Body request
	err := c.Bind(&newUSer)

	if err != nil {
		return c.String(http.StatusBadRequest, "Error to create user")
	}

	users = append(users, newUSer)

	return c.JSON(http.StatusCreated, users)
}