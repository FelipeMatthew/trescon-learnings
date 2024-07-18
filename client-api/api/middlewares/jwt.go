package middlewares

import (
	"github.com/felipematthew/go-learnings/client-api/api/auth"
	"github.com/felipematthew/go-learnings/client-api/api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.JWTConfig{
	Claims:     &models.JwtCustomClaims{},
	SigningKey: auth.JwtSecret,
}

func JwtWithConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.JWTWithConfig(config)(next)
}
