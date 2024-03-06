package helpers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseDate(date *any) {
}

func GenerateNewJwtToken(payload jwt.Claims) (string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	privateKey, err := os.ReadFile("/home/sistemas/Escritorio/goApiBookstore/helpers/demo.rsa")
	if err != nil {
		return "", fmt.Errorf("error reading private key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
	}

	return token.SignedString(key)
}

