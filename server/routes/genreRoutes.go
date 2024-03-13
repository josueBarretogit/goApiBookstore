package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func SetupRoutesGenre(r *gin.Engine) {
	genreRoutes := r.Group(consts.GenreModelName)

	genreController := controllers.NewGenreController()
	store := persistence.NewInMemoryStore(time.Second)
	{
		genreRoutes.GET("/list", cache.CachePage(store, time.Minute, genreController.GetGenres()))
		genreRoutes.GET("/:id/books", genreController.GetBookByGenre())
	}
}
