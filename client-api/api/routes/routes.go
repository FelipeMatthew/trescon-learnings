package routes

import (
	"github.com/felipematthew/go-learnings/client-api/api/handlers"
	"github.com/labstack/echo/v4"
)

func Generate(e *echo.Echo) {

	// Admin routes
	admin := e.Group("/admin")

	admin.GET("/", handlers.GetAdmin)
	admin.POST("/", handlers.CreateAdmin)
	admin.DELETE("/:id", handlers.DeleteAdmin)

	// Clients routes
	client := e.Group("/client")

	client.GET("/", handlers.GetAllClients)
	client.GET("/:id", handlers.GetByIdClient)
	client.POST("/", handlers.GetAdmin)
	client.DELETE("/:id", handlers.GetAdmin)
	client.PUT("/:id", handlers.GetAdmin)
	client.PATCH("/", handlers.GetAdmin)

	// TODO: Middleware & JWT Claims

}
