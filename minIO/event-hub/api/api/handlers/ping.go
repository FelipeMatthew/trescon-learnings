package handlers

import (
	"net/http"

	"github.com/felipematthew/eventhub-api/api/utils"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	// Minio variables
	minioConn := utils.CheckMinioConnection()
	minioEndpoint := "localhost"
	minioUser := "admin"
	minioPort := "9000"

	// Db variables
	dbConn := utils.CheckDbConnection()
	dbEndpoint := "localhost"
	dbUser := "postgres"
	dbName := "eventhub"
	dbPort := "5432"

	// version and build
	version := "2.3"
	build := "3.0"

	// Output
	output := echo.Map{
		"minio_config": map[string]interface{}{
			"connected": minioConn,
			"endpoint":  minioEndpoint,
			"user":      minioUser,
			"port":      minioPort,
		},
		"db_config": map[string]interface{}{
			"connected": dbConn,
			"name":      dbName,
			"endpoint":  dbEndpoint,
			"user":      dbUser,
			"port":      dbPort,
		},
		"vesion": version,
		"build":  build,
	}

	return c.JSON(http.StatusOK, output)
}
