package main

import (
	"fmt"
	"github.com/FelipeMatthew/go-learnings/src/functions"
	"github.com/FelipeMatthew/go-learnings/src/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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


	// Middlewares - Groups
	g := e.Group("/admin")

	// Can use
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.GET("/main", middlewares.MainAdmin)

	e.Start(":8000")
}
