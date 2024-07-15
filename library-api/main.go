package main

import (
	"github.com/felipematthew/go-learnings/library-api/api/config"
	"github.com/felipematthew/go-learnings/library-api/api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Generate(e)

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e.Logger.Fatal(e.Start(":8080"))

}
