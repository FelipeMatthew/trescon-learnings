package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)



func MainJwt(c echo.Context) error {
	// Getting the claims to jwt token - user
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name: ", claims["name"], "userID: ", claims["jti"])

	return c.String(http.StatusOK, "hello from main jwt token page")
}