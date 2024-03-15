package helpers

import (
	"os"
	"regexp"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ParseDate(date *any) {
}

func GenerateNewJwtToken(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseStringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func IsNumberOrUndefined(num string) bool {
	return regexp.MustCompile("[0-9]|undefined").MatchString(num)
}

// Format expected: yyyy-mm-dd
func IsDateOrUndefined(date string) bool {
	return regexp.MustCompile(`^(?:\d{4}-\d{2}-\d{2}|undefined)$`).MatchString(date)
}
