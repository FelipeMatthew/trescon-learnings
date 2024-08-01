package main

import (
	"log"

	"github.com/felipematthew/eventhub-api/api/config"
	"github.com/felipematthew/eventhub-api/api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.StartRoutes(e)

	config.DatabaseInit()

	log.Fatal(e.Start(":8000"))
}
