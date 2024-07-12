package handlers

import (
	"net/http"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/models"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Users)
}


func CreateUser(c echo.Context) error {
	users := models.Users
	newUSer := models.UserType{}

	err := c.Bind(&newUSer)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error to create user")
	}

	users = append(users, newUSer)

	return c.JSON(http.StatusCreated, users)
}