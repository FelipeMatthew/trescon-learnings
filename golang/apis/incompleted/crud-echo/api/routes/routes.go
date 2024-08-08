package routes

import (
	"github.com/FelipeMatthew/go-learnings/crud-echo/api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	users := e.Group("/users")

	users.GET("/", handlers.GetAllUsers)
	users.GET("/:userId", handlers.GetUserById)       
	users.POST("/", handlers.CreateUser) 
	users.DELETE("/:userId", handlers.DeleteUserById) 
	users.PUT("/:userId", handlers.UpdateByUserId)
	users.PATCH("/:userId", handlers.UpdateByUserId)
}