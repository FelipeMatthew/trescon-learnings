package handlers

import (
	"net/http"

	"github.com/felipematthew/go-learnings/client-api/api/auth"
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
	db := config.DB()
	admin := new(models.Admin)

	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	// Hash user pass
	hashedPassword, err := auth.HashPassword(admin.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to hash password",
		})
	}

	admin.Password = hashedPassword

	if err := db.Create(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create admin",
		})
	}

	return c.JSON(http.StatusOK, admin)
}

func DeleteAdmin(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	admin := new(models.Admin)

	if err := db.Find(&admin, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	if admin.Id == 0 {
		return c.JSON(http.StatusInternalServerError, "id not founded")
	}

	if err := db.Delete(&admin, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]string{
		"data": "admin deleted successfuly",
	}

	return c.JSON(http.StatusOK, response)
}
