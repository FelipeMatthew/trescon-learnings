package main

import "github.com/labstack/echo/v4"

func registerRoutes(e *echo.Echo) {

	users := e.Group("/api/v1/users")

	users.GET("/", getAllUsers)
	users.POST("/", createUser) 

	// users.GET("/:userId", getUserById)       
	// users.DELETE("/:userId", deleteUserById) 
	// users.PUT("/:userId", completeUpdateByUserId)
	// users.PATCH("/:userId", partialUpdateByUserId)
}