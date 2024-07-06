package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Payload data
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func CreateJwtToken() (string, error) {
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

func MainJwt(c echo.Context) error {
	// Getting the claims to jwt token - user
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name: ", claims["name"], "userID: ", claims["jti"])

	return c.String(http.StatusOK, "hello from main jwt token page")
}