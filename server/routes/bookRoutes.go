package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutesBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group(consts.BookModelName)

	rootDir, errDir := consts.GetRootDir()

	if errDir != nil {
		panic(errDir.Error())
	}

	imageController := controllers.NewBookImageController(rootDir + "/public" + "/images/")

	{
		bookRoutes.POST(consts.RouteBookImageUpload, middleware.VerifyJwt(), imageController.UploadMultipleImageHadler)
	}
}
