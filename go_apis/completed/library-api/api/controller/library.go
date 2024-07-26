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
		Name:   l.Name,
		Author: l.Author,
		Genre:  l.Genre,
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

func UpdateLibrary(c echo.Context) error {
	id := c.Param("id")
	l := new(model.Library)
	db := config.DB()

	// Bind JSON body to the library struct
	if err := c.Bind(l); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Find the library record by ID
	updatedLibrary := new(model.Library)
	if err := db.First(&updatedLibrary, "id = ?", id).Error; err != nil {
		data := map[string]interface{}{
			"message": "record not found",
		}
		return c.JSON(http.StatusNotFound, data)
	}

	// Update the fields with the new data
	updatedLibrary.Name = l.Name
	updatedLibrary.Author = l.Author
	updatedLibrary.Genre = l.Genre

	// Save the updated record to the database
	if err := db.Model(&model.Library{}).Where("id = ?", id).Updates(updatedLibrary).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": updatedLibrary,
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllLibrary(c echo.Context) error {
	db := config.DB()

	libraries := []*model.Library{}

	if err := db.Find(&libraries).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": libraries,
	}

	return c.JSON(http.StatusOK, response)
}

func GetByIdLibrary(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	libraries := []*model.Library{}

	if err := db.Where("id = ?", id).Find(&libraries).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": libraries[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteLibrary(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()
	l := new(model.Library)

	if err := db.Where("id = ?", id).Delete(&l).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"messge": "book deleted successfully",

	}

	return c.JSON(http.StatusOK, response)
}