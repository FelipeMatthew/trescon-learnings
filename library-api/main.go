package main

import (
	"net/http"

	"github.com/felipematthew/go-learnings/library-api/api/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Init with my own echo")
	})

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e.Logger.Fatal(e.Start(":8080"))

}
