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
	ID   uint   `json:"ID"`
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

type BookByGenreDTO struct {
	BookId         uint    `json:"ID"`
	Title          string  `json:"title,omitempty"`
	CoverPhotoUrl  string  `json:"coverPhotoUrl,omitempty"`
	HardCoverPrice float64 `json:"hardCoverPrice"`
	DigitalPrice   float64 `json:"digitalPrice"`
	AudioPrice     float64 `json:"audioPrice"`
}

func (controller *GenreController) GetBookByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookByGenre []BookByGenreDTO

		id := ctx.Param("id")

		err := database.DB.Table("books").
			Select(`books.id as book_id, books.title as book_title, books.title as title, books.cover_photo_url as cover_photo_url, 
			hard_cover_formats.price as hard_cover_price , digital_formats.price as digital_price, audio_book_formats.price as audio_price`).
			Joins("INNER JOIN hard_cover_formats ON hard_cover_formats.book_id = books.id").
			Joins("INNER JOIN digital_formats ON digital_formats.book_id = books.id").
			Joins("INNER JOIN audio_book_formats ON audio_book_formats.book_id = books.id").
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

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.GenreModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": bookByGenre,
		})
	}
}
