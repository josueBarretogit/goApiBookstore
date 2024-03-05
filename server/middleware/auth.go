package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)




func VerifyJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		receivedToken := ctx.Request.Header.Get("authorization")
		token , err = jwt.ParseWithClaims()
	}
}

