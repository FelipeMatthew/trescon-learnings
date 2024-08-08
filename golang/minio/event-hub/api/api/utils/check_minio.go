package utils

import (
	"context"
	"log"

	"github.com/felipematthew/eventhub-api/api/config"
)

func CheckMinioConnection() bool {
	_, err := config.MinioClient.ListBuckets(context.Background())
	if err != nil {
		log.Printf("Error to connect to minIO: %s", err)
		return false
	}
	return true
}
