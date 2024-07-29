package routes

import (
	"pdfstore/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	files := e.Group("/files")

	files.GET("/", handlers.GetFiles)
	files.POST("/", handlers.InsertFiles)
	files.DELETE("/:filename", handlers.DeleteFiles)

	buckets := e.Group("/buckets")

	buckets.GET("/", handlers.GetBuckets)
	buckets.POST("/:bucketName", handlers.CreateBucket)
}
