package handlers

import (
	"context"
	"net/http"
	"pdfstore/api/config"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

func GetBuckets(c echo.Context) error {
	buckets, err := config.MinioClient.ListBuckets(context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"messege": "buckets founded successfuly",
		"buckets": buckets,
	})
}

func CreateBucket(c echo.Context) error {
	newBucket := c.Param("bucketName")

	err := config.MinioClient.MakeBucket(context.Background(), newBucket, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		// Check if bucket already exist
		exist, errBucketExist := config.MinioClient.BucketExists(context.Background(), newBucket)
		if errBucketExist == nil && exist {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Bucket already exists",
				"bucket":  newBucket,
			})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"messege": "bucket created successfuly",
		"info":    newBucket,
	})
}

func DeleteBucket(c echo.Context) error {
	bucketName := c.Param("bucketName")

	err := config.MinioClient.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"messege": "bucket deleted successfuly"})
}
