package auth

import (
	"fmt"
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

	// Criando o token JWT usando a mesma chave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Assinando o token usando JwtSecret
	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	// Log para verificar o token gerado
	fmt.Println("Generated Token:", signedToken)

	return signedToken, nil
}

