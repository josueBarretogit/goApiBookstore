package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutesBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group(consts.BookModelName)


	imageController := controllers.NewImageController( "book")
	bookController := controllers.NewBookController()

	{
		bookRoutes.POST(consts.RouteBookImageUpload, middleware.VerifyImages,  imageController.UploadMultipleImageHandler)
		bookRoutes.PUT("/assignAuthor/:id", bookController.AssignAuthor())
	}
}
