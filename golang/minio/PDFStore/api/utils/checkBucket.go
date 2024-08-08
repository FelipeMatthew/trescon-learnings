package utils

import (
	"context"
	"pdfstore/api/config"
)

func CheckBucketExist(ctx context.Context, bucketName string) error {
	// Check if the bucket exists
	exists, err := config.MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return err
	}
	return nil
}
