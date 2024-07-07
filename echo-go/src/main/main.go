package main

import (
	"fmt"

	"github.com/FelipeMatthew/go-learnings/src/functions"
	"github.com/FelipeMatthew/go-learnings/src/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Working w echo")

	e := echo.New()

	e.Use(middlewares.ServerHeader)
	e.Use(middleware.Static("./"))

	e.GET("/start", functions.GettingStart)
	e.GET("login", functions.Login)
	e.GET("/person/:data", functions.GetPerson)

	e.POST("/car", functions.AddCar)
	e.POST("/moto", functions.AddMoto)
	e.POST("/dog", functions.AddDog)

	// Middlewares
	adminGroup := e.Group("/admin")

	// Can use
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "felpola" && password == "123" {
				return true, nil
		}
		return false, nil
	}))

	adminGroup.GET("/main", middlewares.MainAdmin)

	// Cookies
	cookieGroup := e.Group("/cookie")

	cookieGroup.Use(middlewares.CheckCookie)

	cookieGroup.GET("/main", middlewares.MainCookie)

	// JWT 

	jwtGroup := e.Group("/jwt")

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey: []byte("mySecret"),
		TokenLookup: "cookie:JWTCookie",
	}))

	jwtGroup.GET("/main", middlewares.MainJwt)

	e.Start(":8000")
}
