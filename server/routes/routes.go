package routes

import (
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(model string, controller controllers.IController, router *gin.Engine) {

	group := router.Group(model)
	{
		group.GET("/findall", controller.FindAll)
		group.GET("/findby/:id", controller.FindOneBy)
		group.POST("/save", controller.Create)
		group.PUT("/update", controller.Update)
		group.DELETE("/delete", controller.Delete)
	}

}
