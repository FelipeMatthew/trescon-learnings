package handlers

import (
	"context"
	"net/http"

	"pdfstore/api/config"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

func GetFiles(c echo.Context) error {
	objects := config.MinioClient.ListObjects(context.Background(), "mynewbucket", minio.ListObjectsOptions{Recursive: true})
	fileNames := []string{}

	for object := range objects {
		if object.Err != nil {
			return object.Err
		}
		fileNames = append(fileNames, object.Key)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filenames": fileNames,
	})
}

// TODO: Install files
func DownloadFiles(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func InsertFiles(c echo.Context) error {
	// READ FROM BODY - form-data file
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Open the files
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	// Create an object in MinIO
	bucketName := "mynewbucket"
	objectName := file.Filename
	contentType := file.Header.Get("Content-Type")

	if contentType != "application/pdf" {
		return c.JSON(http.StatusBadRequest, "Invalid type of file")
	}

	// Uploading the file
	info, err := config.MinioClient.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "file uploaded successfuly",
		"info":    info,
	})
}
