package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutesHardCoverBookFormat(r *gin.Engine) {
	hardCoverFormatRoutes := r.Group(consts.HardcoverFormatModelName)
	hardCoverFormatController := controllers.NewHardCoverFormatController()
	{
		hardCoverFormatRoutes.GET("/details/:id", hardCoverFormatController.GetHardCoverFormatDetails())
	}
}
