package middlewares

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "Hi from cookies context")
}

// Communicating w/ login function
func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
			cookie, err := c.Cookie("sessionID")
			if err != nil {
				log.Println(err)
				return err
			}

			if cookie.Value == "some_string" {
				return next(c)
			}

			return c.String(http.StatusBadRequest, "Check the cookie please")
	}
}