package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		token, err := createJwtToken()
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

// Payload data
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"felipe",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	// Create the token
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
 
	// Sign the token
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
