package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response": "Not authorized",
			})
			return
		}
		verified, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(os.Getenv("PRIVATE_KEY")), nil
		})
		if err != nil {
			panic(err)
		}

		claims, ok := verified.Claims.(jwt.MapClaims)

		if !ok && verified.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response": "not authorzed",
			})
			return

		}

		accountID := fmt.Sprintf("%f", claims["accountID"].(float64))

		ctx.Request.Header.Set("username", claims["username"].(string))
		ctx.Request.Header.Set("accountID", accountID)

		ctx.Next()
	}
}
