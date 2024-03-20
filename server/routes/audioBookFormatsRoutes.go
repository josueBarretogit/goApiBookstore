package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutesAudioBookFormat(r *gin.Engine) {
	audioBookFormatRoutes := r.Group(consts.AudioBookFormatModelName)
	accountController := controllers.NewAudioBookFormatController()
	{
		audioBookFormatRoutes.GET("/details/:id", accountController.GetAudioDetails())
	}
}
