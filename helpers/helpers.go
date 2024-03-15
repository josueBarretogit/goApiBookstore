package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SQLBuilder struct {
	selectSentence    string
	tableName         string
	leftJoinsSentence string
	innerJoins        string
	PreparedSentence  string
}

func (builder *SQLBuilder) Select(columns ...string) *SQLBuilder {
	builder.selectSentence = `SELECT `
	for _, column := range columns {
		builder.selectSentence += fmt.Sprintf(`%s,`, column)
	}
	builder.selectSentence = strings.Trim(builder.selectSentence, ",")
	return builder
}

func (builder *SQLBuilder) LeftJoins(join string) *SQLBuilder {
	builder.leftJoinsSentence += fmt.Sprintf(" LEFT JOIN %s ", join)
	return builder
}

func (builder *SQLBuilder) InnerJoins(join string) *SQLBuilder {
	builder.innerJoins += fmt.Sprintf(" INNER JOIN %s ", join)
	return builder
}

func (builder *SQLBuilder) GetSQL() string {
	builder.PreparedSentence = builder.selectSentence + fmt.Sprintf("FROM %s ", builder.tableName) + builder.leftJoinsSentence
	return builder.PreparedSentence
}

func NewSQLBuilder(tablename string) *SQLBuilder {
	return &SQLBuilder{}
}

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
