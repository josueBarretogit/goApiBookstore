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
		authorRoutes.GET("/:id/books", authorController.GetAuthorBooks())
		authorRoutes.PUT("/assignPublisher/:id", authorController.AssignPublisher())
		authorRoutes.PUT("/assignBook/:id", authorController.AssignBook())
	}
}
