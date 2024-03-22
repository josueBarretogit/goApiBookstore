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
	TableName         string
	leftJoinsSentence string
	innerJoins        string
	whereSentence     string
	groupBy           string
	pagination        string
	orderBy           string
	PreparedSentence  string
}

func (builder *SQLBuilder) Select(columns ...string) *SQLBuilder {
	builder.selectSentence = `SELECT ` + strings.Join(columns, ",")

	return builder
}

func (builder *SQLBuilder) LeftJoins(tableToJoin, condition string) *SQLBuilder {
	builder.leftJoinsSentence += fmt.Sprintf(" LEFT JOIN %s ON %s ", tableToJoin, condition)
	return builder
}

func (builder *SQLBuilder) InnerJoins(tableToInnerJoin string, condition string) *SQLBuilder {
	builder.innerJoins += fmt.Sprintf(" INNER JOIN %s ON %s ", tableToInnerJoin, condition)
	return builder
}

func (builder *SQLBuilder) Where(condition string) *SQLBuilder {
	builder.whereSentence = fmt.Sprintf(` WHERE %s`, condition)
	return builder
}

func (builder *SQLBuilder) AndWhere(condition string) *SQLBuilder {
	builder.whereSentence += fmt.Sprintf(` AND %s `, condition)
	return builder
}

func (builder *SQLBuilder) OrWhere(condition string) *SQLBuilder {
	builder.whereSentence += fmt.Sprintf(` OR %s`, condition)
	return builder
}

func (builder *SQLBuilder) Group() *SQLBuilder {
	builder.groupBy = ` GROUP BY `
	return builder
}

func (builder *SQLBuilder) BY(groupBy string) *SQLBuilder {
	builder.groupBy += ` ` + groupBy
	return builder
}

func (builder *SQLBuilder) Paginate(page int, itemsPerPage int) *SQLBuilder {
	builder.pagination = fmt.Sprintf(` LIMIT %d OFFSET %d`, itemsPerPage, (page-1)*itemsPerPage)
	return builder
}

func (builder *SQLBuilder) OrderBy(conditions string) *SQLBuilder {
	builder.orderBy = fmt.Sprintf(` ORDER BY  %s`, conditions)
	return builder
}

func (builder *SQLBuilder) GetSQL() string {
	builder.PreparedSentence = builder.selectSentence +
		fmt.Sprintf(" FROM %s ", builder.TableName) +
		builder.leftJoinsSentence +
		builder.innerJoins +
		builder.whereSentence +
		builder.groupBy +
		builder.orderBy +
		builder.pagination

	return builder.PreparedSentence
}

func NewSQLBuilder(tablename string) *SQLBuilder {
	return &SQLBuilder{
		TableName: tablename,
	}
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

func IsNumber(num string) bool {
	pattern := `^-?\d+$`
	return regexp.MustCompile(pattern).MatchString(num)
}

// Format expected: yyyy-mm-dd
func IsDateOrUndefined(date string) bool {
	return regexp.MustCompile(`^(?:\d{4}-\d{2}-\d{2}|undefined)$`).MatchString(date)
}
