package functions

import (
	"log"
	"net/http"
	"time"

	"github.com/FelipeMatthew/go-learnings/src/middlewares"
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

		// JWT token
		token, err := middlewares.CreateJwtToken()
		if err != nil {
			log.Println("error creating the jwt token", err)
			return c.String(http.StatusInternalServerError, "Error to create jwt token")
		}

		// JWT cookie 
		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusInternalServerError, "Invalid credentials, try again")
}
