package handlers

import (
	"context"
	"net/http"
	"pdfstore/api/config"

	"github.com/labstack/echo/v4"
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
