package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"
	"api/bookstoreApi/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutesBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group(consts.BookModelName)

	rootDir, errDir := consts.GetRootDir()

	if errDir != nil {
		panic(errDir.Error())
	}

	imageService :=  services.NewImageService(50)

	imageController := controllers.NewBookImageController(rootDir + "/public" + "/images/", imageService)
	bookController := controllers.NewBookController()

	{
		bookRoutes.POST(consts.RouteBookImageUpload, middleware.VerifyJwt(), imageController.UploadMultipleImageHandler)
		bookRoutes.PUT("/assignAuthor/:id", bookController.AssignAuthor())
	}
}
