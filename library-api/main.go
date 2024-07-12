package main

import (
	"github.com/felipematthew/go-learnings/library-api/api/config"
	"github.com/felipematthew/go-learnings/library-api/api/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", controller.GetLibrary)
	e.POST("/", controller.CreateLibrary)

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e.Logger.Fatal(e.Start(":8080"))

}
