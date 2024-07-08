package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetPerson(c echo.Context) error {
	personName := c.QueryParam("name")
	personType := c.QueryParam("type")
	
	// Must be the same passed on e.get
	dataType := c.Param("data")

	if dataType == 	"string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your person name is: %s\nand his type is: %s\n", personName, personType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": personName,
			"type": personType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "invalid data type",
	})
}