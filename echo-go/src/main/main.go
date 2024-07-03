package main

import (
	"fmt"
	"github.com/FelipeMatthew/go-learnings/src/functions"

	"github.com/labstack/echo"
)



func main() {
	fmt.Println("Working w echo")

	e := echo.New()

	e.GET("/", functions.GettingStart)
	e.GET("/person/:data", functions.GetPerson)

	// Requests
	e.POST("/car", functions.AddCar)
	e.POST("/moto", functions.AddMoto)
	e.POST("/dog", functions.AddDog)

	e.Start(":8000")
}
