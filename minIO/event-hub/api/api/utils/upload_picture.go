package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/felipematthew/eventhub-api/api/config"
	"github.com/minio/minio-go/v7"
)

// UploadPicture uploads a picture to MinIO and returns the presigned URL
func UploadPicture(file *multipart.FileHeader) (string, error) {
	bucketName := "eventhub"

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Generate an object name based on the current timestamp and file name
	objectName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	contentType := file.Header.Get("Content-Type")

	// Upload the file to MinIO
	_, err = config.MinioClient.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	// Generate a presigned URL valid for 7 days
	reqParams := make(url.Values)
	presignedURL, err := config.MinioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Hour*24*7, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
