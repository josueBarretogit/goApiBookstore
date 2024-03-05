package helpers

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseDate(date *any) {
}

func GenerateNewJwtToken(payload jwt.Claims) (string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	return token.SignedString(os.Getenv("SECRET_KEY"))
}
