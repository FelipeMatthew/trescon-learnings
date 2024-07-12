package main

import (
	"github.com/FelipeMatthew/go-learnings/api-go/cmd/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers. PostSingleHandler)

	e.Logger.Fatal(e.Start(":8888"))
}