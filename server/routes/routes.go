package routes

import (
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(model string, controller controllers.IController, router *gin.Engine) {
	group := router.Group(model)
	{
		group.GET("/findall", controller.FindAll())
		group.GET("/findby/:id", middleware.VerifyJwt() , controller.FindOneBy())
		group.POST("/save", controller.Create())
		group.PUT("/update/:id", controller.Update())
		group.DELETE("/delete/:id", controller.Delete())
	}
}
