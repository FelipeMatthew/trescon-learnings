package middlewares

import "github.com/labstack/echo"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "FelpsServer/1.0")
		c.Response().Header().Set("my-header-content", "FelpolaaServer/1.0")
		return next(c)
	}
}
