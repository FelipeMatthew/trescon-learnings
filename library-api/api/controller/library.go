package controller

import (
	"net/http"

	"github.com/felipematthew/go-learnings/library-api/api/config"
	"github.com/felipematthew/go-learnings/library-api/api/model"
	"github.com/labstack/echo/v4"
)

func CreateLibrary(c echo.Context) error {
	l := new(model.Library)
	db := config.DB()

	// Getting data
	if err := c.Bind(l); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	libary := &model.Library{
		Name: l.Name,
		Author: l.Author,
		Genre: l.Genre,
	}

	// DB Create 
	if err := db.Create(&libary).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": l,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetLibrary(c echo.Context) error {
	return c.String(200, "testing")
}