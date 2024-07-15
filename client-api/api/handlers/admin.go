package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "getting admin account")
}

func CreateAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "getting admin account")
}

func DeleteAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "getting admin account")
}