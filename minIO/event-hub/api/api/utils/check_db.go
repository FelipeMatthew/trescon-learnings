package utils

import (
	"context"
	"log"

	"github.com/felipematthew/eventhub-api/api/config"
)

func CheckDbConnection() bool {
	sqlDb, err := config.DB().DB()
	if err != nil {
		log.Printf("Failed to get database handle: %v", err)
		return false
	}

	err = sqlDb.PingContext(context.Background())
	if err != nil {
		log.Printf("Database connection error: %v", err)
		return false
	}

	return true
}
