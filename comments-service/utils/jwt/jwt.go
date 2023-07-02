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

func GenerateToken(userID int) (string, error) {
	// Define la fecha de expiración del token
	expirationTime := time.Now().Add(24 * time.Hour) // Ejemplo: expira en 24 horas

	// Crea los claims (datos) del token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Id:        fmt.Sprintf("%d", userID),
	}

	// Crea el token JWT utilizando la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
