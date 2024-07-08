package router

import (
	"github.com/FelipeMatthew/go-learnings/src/api"
	"github.com/FelipeMatthew/go-learnings/src/api/middlewares"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Groups
	cookieGroup := e.Group("/cookie")
	adminGroup := e.Group("/admin")
	jwtGroup := e.Group("/jwt")

	// Set Middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddleware(cookieGroup)
	middlewares.SetJwtMiddleware(jwtGroup)

	// Main Routes
	api.MainGroup(e)

	// Groups Routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}