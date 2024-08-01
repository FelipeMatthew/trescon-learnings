package routes

import (
	"github.com/felipematthew/eventhub-api/api/handlers"
	"github.com/labstack/echo/v4"
)

func StartRoutes(e *echo.Echo) {

	admin := e.Group("/admin")

	admin.GET("/", handlers.GetAdmins)

}
