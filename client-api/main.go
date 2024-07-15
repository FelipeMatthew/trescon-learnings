package main

import (
	"log"

	"github.com/felipematthew/go-learnings/client-api/api/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error  {
		return c.String(200, "hiii")
	})

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	log.Fatal(e.Start(":8080"))
}