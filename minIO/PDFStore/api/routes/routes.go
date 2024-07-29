package routes

import (
	"pdfstore/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	minio := e.Group("/minio")

	minio.GET("/", handlers.GetFiles)
	minio.POST("/", handlers.InsertFiles)
}
