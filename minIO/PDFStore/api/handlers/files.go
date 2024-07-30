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
	folderPath := c.Param("*") // Geting all or nothing
	file, err := c.FormFile("file")

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	defer src.Close()

	// Validating folder path
	var objectName string

	switch folderPath {
	case "":
		objectName = file.Filename
	case "documentations":
		objectName = "documentations/" + file.Filename
	case "certifications":
		objectName = "certifications/" + file.Filename
	default:
		return c.JSON(http.StatusBadRequest, "Folder not found")
	}

	// Check content type
	contentType := file.Header.Get("Content-Type")
	if contentType != "application/pdf" {
		return c.JSON(http.StatusBadRequest, "Invalid type of file")
	}

	// Check if the bucket exists
	exists, err := config.MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to check if bucket exists: " + err.Error(),
		})
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Bucket does not exist",
		})
	}

	// Uploading the file
	info, err := config.MinioClient.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to upload file: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "File uploaded successfully",
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
