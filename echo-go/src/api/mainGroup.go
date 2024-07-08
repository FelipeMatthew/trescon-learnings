package api

import (
	"github.com/FelipeMatthew/go-learnings/src/api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/start", handlers.GettingStart)
	e.GET("/login", handlers.Login)
	e.GET("/person/:data", handlers.GetPerson)

	e.POST("/car", handlers.AddCar)
	e.POST("/moto", handlers.AddMoto)
	e.POST("/dog", handlers.AddDog)
}