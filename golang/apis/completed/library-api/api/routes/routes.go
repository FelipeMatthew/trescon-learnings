package routes

import (
	"github.com/felipematthew/go-learnings/library-api/api/controller"
	"github.com/labstack/echo/v4"
)

func Generate(e *echo.Echo) *echo.Echo {
	
	library := e.Group("/library")

	library.GET("/", controller.GetAllLibrary)
	library.POST("/", controller.CreateLibrary)
	library.GET("/:id", controller.GetByIdLibrary)
	library.PUT("/:id", controller.UpdateLibrary)
	library.DELETE("/:id", controller.DeleteLibrary)

	return e
}