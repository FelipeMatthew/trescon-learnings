package main

import (
	"pdfstore/api/config"
	"pdfstore/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.MinioInit()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
