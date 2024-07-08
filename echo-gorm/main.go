package main

import (

	//	"gorm.io/driver/postgres"
	"net/http"

	"github.com/FelipeMatthew/go-learnings/echo-gorm/controller"
	"github.com/FelipeMatthew/go-learnings/echo-gorm/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/students", controller.GetStudents)

	storage.NewDB()

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello from main page")
}