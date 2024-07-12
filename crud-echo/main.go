package main

import (
	"fmt"

	"github.com/FelipeMatthew/go-learnings/crud-echo/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New() // echo instance

	// route
	routes.RegisterRoutes(e)
	fmt.Println("Hello Echo")

	// middleware
	e.Use(middleware.Logger())

	// run app
	e.Start(":8080")
}

