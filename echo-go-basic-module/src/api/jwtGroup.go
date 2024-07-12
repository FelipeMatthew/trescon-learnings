package api

import (
	"github.com/FelipeMatthew/go-learnings/src/api/handlers"
	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}