package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(model string, controller controllers.IController, router *gin.Engine) {
	store := persistence.NewInMemoryStore(time.Second)
	group := router.Group(model)
	{
		group.GET(consts.RouteFindAll, middleware.VerifyJwt(), cache.CachePage(store, time.Minute, controller.FindAll()))
		group.GET(consts.RouteFindById, middleware.VerifyJwt(), cache.CachePage(store, time.Minute, controller.FindOneBy()))
		group.POST(consts.RouteCreate, controller.Create())
		group.PUT(consts.RouteUpdate, controller.Update())
		group.DELETE(consts.RouteDelete, controller.Delete())

	}
}
