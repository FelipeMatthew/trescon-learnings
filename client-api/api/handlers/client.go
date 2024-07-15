package handlers

import (
	"net/http"

	"github.com/felipematthew/go-learnings/client-api/api/config"
	"github.com/felipematthew/go-learnings/client-api/api/models"
	"github.com/labstack/echo/v4"
)

func GetAllClients(c echo.Context) error {
	db := config.DB()
	clients := []*models.Clients{}

	if err := db.Find(&clients).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"status": 200,
		"data":   clients,
	}

	return c.JSON(http.StatusOK, response)
}
