package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutesBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group(consts.BookModelName)

	imageController := controllers.NewImageController("book")
	bookController := controllers.NewBookController()

	{

		bookRoutes.GET(consts.RouteBestSellers+"/page/:page/itemsPerPage/:itemsPerPage/genre/:idGenre", bookController.GetBestSellers())
		bookRoutes.GET("/formats/:id", bookController.GetBookFormats())
		bookRoutes.GET("/test", bookController.Test())
		bookRoutes.GET("/reviews/:id/page/:page/itemsPerPages/:itemsPerPages/rating/:rating", bookController.GetReviews())
		bookRoutes.GET("/one/:idBook", bookController.GetOneBook())
		bookRoutes.GET("/reviewStatistics/:idBook", bookController.GetReviewStatistics())
		bookRoutes.GET(`/search/:searchTerm/page/:page/itemsPerPage/:itemsPerPage/genre/:genre/rangePrice/:rangePrice1/:rangePrice2/highToLowPrice/:highToLowPrice/rating/:rating/fromDate/:fromDate/toDate/:toDate/language/:idLanguage`, bookController.SearchBook())
		bookRoutes.POST(consts.RouteBookImageUpload, middleware.VerifyImages, imageController.UploadMultipleImageHandler)
		bookRoutes.PUT("/assignAuthor/:id", bookController.AssignAuthor())
		bookRoutes.PUT("/assignLanguage/:id", bookController.AssignLanguage())
		bookRoutes.PUT("/:id/rating/:rating", bookController.UpdateRating())
	}
}
