package routes

import (
	"pdfstore/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	files := e.Group("/files")

	files.GET("/:bucketName", handlers.GetFiles)
	files.GET("/:bucketName/download", handlers.DownloadFiles)
	files.GET("/:bucketName/:filename/temp-url", handlers.GenerateTemporaryUrl)
	files.POST("/:bucketName/*", handlers.InsertFiles)
	files.DELETE("/:bucketName/:filename", handlers.DeleteFiles)

	buckets := e.Group("/buckets")

	buckets.GET("/", handlers.GetBuckets)
	buckets.POST("/:bucketName", handlers.CreateBucket)
	buckets.DELETE("/:bucketName", handlers.DeleteBucket)
}
