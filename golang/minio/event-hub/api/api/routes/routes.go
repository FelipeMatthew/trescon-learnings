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

	// Admin
	admin.GET("/", handlers.GetAdmins)
	admin.POST("/", handlers.CreateAdmin)
	admin.POST("/login", handlers.Login)
	admin.DELETE("/:id", handlers.DeleteAdmin)

	// Event
	event.POST("/", handlers.GetAdmins)

	// Participant
	participant.POST("/", handlers.CreateParticipant)

}
