package routes

import (
	"github.com/felipematthew/go-learnings/client-api/api/controllers"
	"github.com/labstack/echo/v4"
)

func Generate(e *echo.Echo) {

	// Admin routes
	admin := e.Group("/admin")

	admin.GET("/", controllers.GetAdmin)
	admin.POST("/", controllers.CreateAdmin)
	admin.DELETE("/:id", controllers.DeleteAdmin)

	// Clients routes
	client := e.Group("/client")

	client.GET("/", controllers.GetAdmin)
	client.POST("/", controllers.GetAdmin)
	client.GET("/:id", controllers.GetAdmin)
	client.DELETE("/:id", controllers.GetAdmin)
	client.PUT("/:id", controllers.GetAdmin)
	client.PATCH("/", controllers.GetAdmin)

	// TODO: Middleware & JWT Claims

}
