package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutesPublisher(r *gin.Engine) {
	publisherRoutes := r.Group(consts.AuthorModelName)
	publisherController := controllers.NewPublisherController()
	{
		publisherRoutes.PUT("/assignAuthor/:id", publisherController.AssignAuthor())
	}
}
