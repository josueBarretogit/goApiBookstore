package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func VerifyImages(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		switch file.Header.Get("Content-Type") {
		case "image/jpeg" , "image/png" , "image/webp":
		default:
			ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
				"response" : "that type of file is not allowed",
				"success" : false,
			})
			return 
		} 
	}
	ctx.Next()
	
}
