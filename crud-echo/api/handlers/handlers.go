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

	// Body request
	err := c.Bind(&newUSer)

	if err != nil {
		return c.String(http.StatusBadRequest, "Error to create user")
	}

	users = append(users, newUSer)

	return c.JSON(http.StatusCreated, users)
}

func GetUserById(c echo.Context) error {
	userId := c.Param("userId")

	for _, user := range models.Users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.String(http.StatusBadRequest, "User not found")
}

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