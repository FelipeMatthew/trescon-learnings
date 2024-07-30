package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"pdfstore/api/config"
	"pdfstore/api/utils"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

func GetFiles(c echo.Context) error {
	bucketName := c.Param("bucketName")

	// Check if the bucket exists
	if err := utils.CheckBucketExist(context.Background(), bucketName); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to check the bucket: " + err.Error(),
		})
	}

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

func DownloadFiles(c echo.Context) error {
	bucketName := c.Param("bucketName")
	localDir := "./files/" + bucketName

	// Create dir if dont exist
	// Full folder access
	if err := os.Mkdir(localDir, os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create local directory: %v", err))
	}

	// Get all objects
	objects := config.MinioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true})

	// Get one by one
	for object := range objects {
		// Validating
		if object.Err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error listing object: %v", object.Err))
		}

		// Getting the object name
		objectName := object.Key
		// ./files/objectname
		localFilePath := filepath.Join(localDir, objectName) // Vai unir os caracteres com base no OS

		// Criar diretórios locais necessários para o arquivo
		if err := os.MkdirAll(filepath.Dir(localFilePath), os.ModePerm); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create local directories: %v", err))
		}

		// Abrir novo arquivo local para salva o conteudo
		localFile, err := os.Create(localFilePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create local file path: %v", err))
		}
		// Vai fechar no final do loop
		defer localFile.Close()

		// * Apenas pegar o objeto do MinIO
		objectReader, err := config.MinioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to get object from MinIO: %v", err))
		}
		defer objectReader.Close()

		// * copiar o objeto do MinIO para maquina local
		if _, err := io.Copy(localFile, objectReader); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to copy object to local file: %v", err))
		}
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("All files downloaded successfully to %s", localDir))
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
	if err := utils.CheckBucketExist(context.Background(), bucketName); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to check the bucket: " + err.Error(),
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

func GenerateTemporaryUrl(c echo.Context) error {
	bucketName := c.Param("bucketName")
	fileName := c.Param("filename")

	temporaryTime := time.Second * 24 * 60 * 60 // One day

	reqParms := make(url.Values)
	presignedUrl, err := config.MinioClient.PresignedGetObject(context.Background(), bucketName, fileName, temporaryTime, reqParms)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"url": presignedUrl.String(),
	})
}
