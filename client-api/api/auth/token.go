package auth

import (
	"time"

	"github.com/felipematthew/go-learnings/client-api/api/models"
	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte("your_secret_key")

func GenerateJwt(name string, admin bool) (string, error) {
	claims := &models.JwtCustomClaims{
		Name:  name,
		Admin: admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Getting all config and hasing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Getting the -> your_secret_key -> and signed to string
	return token.SignedString(JwtSecret)
}
