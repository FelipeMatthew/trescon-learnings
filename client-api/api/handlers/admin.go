package handlers

import (
	"net/http"

	"github.com/felipematthew/go-learnings/client-api/api/config"
	"github.com/felipematthew/go-learnings/client-api/api/models"
	"github.com/labstack/echo/v4"
)

func GetAdmin(c echo.Context) error {
	db := config.DB()

	admins := []*models.Admin{}

	if err := db.Find(&admins).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": admins, 
	}

	return c.JSON(http.StatusOK, response)
}

func CreateAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "getting admin account")
}

func DeleteAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "getting admin account")
}