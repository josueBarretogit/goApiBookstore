package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutesDigitlBookFormat(r *gin.Engine) {
	digitalFormatRoutes := r.Group(consts.DigitalFormatModelName)
	digitalFormatController := controllers.NewDigitalFormatController()
	{
		digitalFormatRoutes.GET("/details/:id", digitalFormatController.GetDigitalFormatDetails())
	}
}
