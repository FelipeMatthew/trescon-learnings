package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}


func createUser(c echo.Context) error {
	var newUser userType

	err := c.Bind(&newUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	users = append(users, newUser)
	return c.JSON(http.StatusCreated, newUser)
}

// TODO
// func getUserById(){}

// TODO
// func deleteUserById(){}

// TODO
// func completeUpdateByUserId(){}

// TODO
// func partialUpdateByUserId(){}

// TODO
// func findUserById() {}