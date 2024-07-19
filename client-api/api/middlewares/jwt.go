package middlewares

import (
	"fmt"

	"github.com/felipematthew/go-learnings/client-api/api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.JWTConfig{
	Claims:        &models.JwtCustomClaims{},
	SigningMethod: "HS512",
	SigningKey:    []byte("your_secret_key"),
}

func JwtWithConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Middleware JWT Called")
		return middleware.JWTWithConfig(config)(next)(c)
	}
}
