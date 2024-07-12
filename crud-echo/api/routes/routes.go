package routes

import (
	"github.com/FelipeMatthew/go-learnings/crud-echo/api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	users := e.Group("/users")

	users.GET("/", handlers.GetAllUsers)
	users.POST("/", handlers.CreateUser) 

	// users.GET("/:userId", getUserById)       
	// users.DELETE("/:userId", deleteUserById) 
	// users.PUT("/:userId", completeUpdateByUserId)
	// users.PATCH("/:userId", partialUpdateByUserId)
}