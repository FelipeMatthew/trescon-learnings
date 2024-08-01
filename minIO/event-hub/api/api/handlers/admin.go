package handlers

import (
	"net/http"

	"github.com/felipematthew/eventhub-api/api/config"
	"github.com/felipematthew/eventhub-api/api/models"
	"github.com/labstack/echo/v4"
)

func GetAdmins(c echo.Context) error {
	db := config.DB()

	admins := []*models.Admin{}

	if err := db.Find(&admins).Error; err != nil {
		data := echo.Map{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := echo.Map{
		"data": admins,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateAdmin(c echo.Context) error {
	db := config.DB()

	username := c.FormValue("username")
	password := c.FormValue("password")

	// Insert on db
	insert := db.Exec("INSERT INTO admins (username, password) VALUES (?, crypt(?, gen_salt('bf')))", username, password)
	if insert.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": insert.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}

func Login(c echo.Context) error {
	db := config.DB()

	username := c.FormValue("username")
	password := c.FormValue("password")

	var count int64
	// Query
	query := "SELECT COUNT(*) FROM admins WHERE username = ? AND password = crypt(?, password)"
	result := db.Raw(query, username, password).Count(&count)

	// Vai pegar e se o contador achar o user e senha ele retorna 1
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	if count > 0 {
		return c.JSON(http.StatusOK, echo.Map{"message": "Authentication successful"})
		// TODO JWT
	}

	return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
}
