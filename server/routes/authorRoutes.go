package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutesAuthor(r *gin.Engine) {
	authorRoutes := r.Group(consts.AuthorModelName)
	authorController := controllers.NewAuthorController()
	{
		authorRoutes.PUT("/assignPublisher/:id", authorController.AssignPublisher())
	}
}
