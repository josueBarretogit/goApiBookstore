package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)




func VerifyJwt() gin.HandlerFunc {
	publicKey , errKey := os.ReadFile("/home/sistemas/Escritorio/goApiBookstore/helpers/public.rsa.pub")
	if errKey != nil {
		panic(errKey.Error())
	}
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("authorization")	
		verified , err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
			if err != nil {
				return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
			}

			return key, nil
		})
		if err != nil {
			panic(err)
		}


		if claims, ok := verified.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["accountId"])
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"response" : "not authorzed",
			})
			return
		}
	}
}

