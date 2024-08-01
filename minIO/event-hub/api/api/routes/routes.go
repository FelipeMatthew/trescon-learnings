package routes

import (
	"github.com/felipematthew/eventhub-api/api/handlers"
	"github.com/labstack/echo/v4"
)

func StartRoutes(e *echo.Echo) {

	admin := e.Group("/admin")
	event := e.Group("/event")
	participant := e.Group("/participant")

	// Ping
	e.GET("/ping", handlers.Ping)

	admin.GET("/", handlers.GetAdmins)
	event.GET("/", handlers.GetAdmins)
	participant.GET("/", handlers.GetAdmins)

}
