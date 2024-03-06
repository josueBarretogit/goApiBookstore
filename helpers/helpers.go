package helpers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
