package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(model string, controller controllers.IController, router *gin.Engine) {
	group := router.Group(model)
	{
		group.GET(consts.RouteFindAll, middleware.VerifyJwt(), controller.FindAll())
		group.GET(consts.RouteFindById, middleware.VerifyJwt(), controller.FindOneBy())
		group.POST(consts.RouteCreate, controller.Create())
		group.PUT(consts.RouteUpdate, controller.Update())
		group.DELETE(consts.RouteDelete, controller.Delete())

	}
}
