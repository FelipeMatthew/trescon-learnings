package main

import (
	"log"
	"pdfstore/api/config"
	"pdfstore/api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	e := echo.New()

	config.MinioInit()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
