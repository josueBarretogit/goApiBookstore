package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutesCustomer(r *gin.Engine) {
	customerRoutes := r.Group(consts.CustomerModelName)

	customerController := controllers.NewCustomerController()
	imageController := controllers.NewImageController("customer")
	{
		customerRoutes.GET("/:id/reviews", customerController.GetCustomerReviews())
		customerRoutes.POST("/uploadCustomerProfilePicture/:id", middleware.VerifyImages, imageController.UploadMultipleImageHandler)
	}
}
