package functions

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// validations
	if username == "felipe" && password == "123" {
		cookie := &http.Cookie{}

		// cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string" // Hashed value
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You are logged in!")
	}

	return c.String(http.StatusOK, "Invalid credentials, try again")
}
