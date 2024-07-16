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

func GetByIdClient(c echo.Context) error {
	db := config.DB()
	id := c.Param("id")
	client := new(models.Clients)

	if err := db.Where("id = ?", id).Find(&client).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	if client.Id == 0 {
		data := map[string]string{
			"message": "user not founded",
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"status": 200,
		"data":   client,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateClient(c echo.Context) error {
	db := config.DB()
	client := &models.Clients{}

	// Body
	if err := c.Bind(&client); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	newClient := &models.Clients{
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Email:     client.Email,
		Age:       client.Age,
		Phone:     client.Phone,
	}

	if err := db.Create(&newClient).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"status": "201",
		"data":   newClient,
	}

	return c.JSON(http.StatusCreated, response)
}

func DeleteClient(c echo.Context) error {
	db := config.DB()
	id := c.Param("id")
	client := new(models.Clients)

	if err := db.Delete(&client, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "user deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}

func CompleteEditClient(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()
	client := new(models.Clients)

	// Body req
	if err := c.Bind(&client); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Get in db by id
	newClient := new(models.Clients)
	if err := db.Find(&newClient, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Updates then
	newClient.FirstName = client.FirstName
	newClient.LastName = client.LastName
	newClient.Email = client.Email
	newClient.Age = client.Age
	newClient.Phone = client.Phone

	// Save in db
	if err := db.Model(&models.Clients{}).Where("id = ?", id).Updates(newClient).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"messaage": "user edited successfuly",
		"data":     newClient,
	}

	return c.JSON(http.StatusOK, response)
}

func PartialEditClient(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()
	client := new(models.Clients)
}
