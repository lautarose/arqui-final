package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey = []byte("secret") //must be a local environment variable.
)

func VerifyToken(authToken string) (*jwt.StandardClaims, error) {
	tokenString := strings.Split(authToken, " ")[1]
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	// Verifica si el token ha expirado
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, fmt.Errorf("el token ha expirado")
	}

	return claims, nil
}
