package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/FelipeMatthew/go-learnings/echo-gorm/model"
	"github.com/FelipeMatthew/go-learnings/echo-gorm/storage"
)

// Handlers
func GetStudents(c echo.Context) error {
	students, _ := 	getRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func getRepoStudents() ([]model.Students, error) {
	db := storage.GetDbInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}