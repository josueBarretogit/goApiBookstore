package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreController struct {
	GenericController[usermodels.Genre]
}

type GetGenresListDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (controller *GenreController) GetGenres() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var listGenres []GetGenresListDTO

		err := database.DB.Model(&usermodels.Genre{}).Select("name, id").Scan(&listGenres)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.GenreModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"list": listGenres,
		})
	}
}

type BookGetAllDTO struct {
	BookId        uint    `json:"id"`
	Title         string  `json:"title,omitempty"`
	CoverPhotoUrl string  `json:"cover_photo_url,omitempty"`
	Price         float64 `json:"price"`
}

type BookByGenreDTO struct {
	Name string `json:"name"`
}

func (controller *GenreController) GetBookByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookByGenre []BookGetAllDTO
		var genre BookByGenreDTO

		id := ctx.Param("id")

		err := database.DB.Table("books").
			Select("books.id as book_id, books.title as book_title, books.title as title, books.cover_photo_url as cover_photo_url, hard_cover_formats.price as price ").
			Joins("INNER JOIN hard_cover_formats ON hard_cover_formats.book_id = books.id").
			Joins("INNER JOIN genres ON genres.id = books.genre_id").
			Where("genres.id = ?", id).
			Scan(&bookByGenre)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.GenreModelName,
			})
			return
		}

		err = database.DB.Model(&usermodels.Genre{}).
			Select("name").
			Where("genres.id = ?", id).
			Find(&genre)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.GenreModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"genre": genre,
			"books": bookByGenre,
		})
	}
}
