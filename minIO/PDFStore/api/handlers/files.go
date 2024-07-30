package handlers

import (
	"context"
	"net/http"

	"pdfstore/api/config"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

// TODO: Fazer pastas e dividir os arquivos entre elas

func GetFiles(c echo.Context) error {
	bucketName := c.Param("bucketName")

	objects := config.MinioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true})
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
	bucketName := c.Param("bucketName")

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

func DeleteFiles(c echo.Context) error {
	bucketName := c.Param("bucketName")
	objectName := c.Param("filename")

	// TODO: Validation with not founded file

	err := config.MinioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	message := "file: " + objectName + " deleted successfuly"

	return c.JSON(http.StatusOK, echo.Map{"message": message})
}
