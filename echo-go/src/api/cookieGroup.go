package api

import (
	"github.com/FelipeMatthew/go-learnings/src/api/handlers"
	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}