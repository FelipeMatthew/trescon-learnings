package routes

import (
	"github.com/felipematthew/go-learnings/client-api/api/handlers"
	"github.com/felipematthew/go-learnings/client-api/api/middlewares"
	"github.com/labstack/echo/v4"
)

func Generate(e *echo.Echo) {

	e.POST("/login", handlers.Login)

	// Admin routes
	admin := e.Group("/admin")

	admin.Use(middlewares.JwtWithConfig)

	admin.GET("/", handlers.GetAdmin)
	admin.POST("/", handlers.CreateAdmin)
	admin.DELETE("/:id", handlers.DeleteAdmin)

	// Clients routes
	client := e.Group("/client")

	client.GET("/", handlers.GetAllClients)
	client.GET("/:id", handlers.GetByIdClient)
	client.POST("/", handlers.CreateClient)
	client.DELETE("/:id", handlers.DeleteClient)
	client.PUT("/:id", handlers.CompleteUpdateClient)
	client.PATCH("/:id", handlers.PartialUpdateClient)

	// TODO: Middleware & JWT Claims

}
