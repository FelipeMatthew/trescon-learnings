package main

import (
	"log"

	"github.com/felipematthew/go-learnings/client-api/api/config"
	"github.com/felipematthew/go-learnings/client-api/api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Generate(e)

	config.DatabaseInit()

	log.Fatal(e.Start(":8080"))
}