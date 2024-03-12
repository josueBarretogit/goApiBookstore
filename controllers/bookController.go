package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	GenericController[usermodels.Book]
}

func (controller *BookController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Book, usermodels.Author]("Authors", consts.BookModelName)
}

type BestSellerBooks struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	CoverPhotoUrl string `json:"cover_photo_url"`
	Rating        *int   `json:"rating,omitempty"`
	TotalSold     int    `json:"total_sold" gorm:"column:total_sold"`
}

func (controller *BookController) GetBestSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mostSelledBooks []BestSellerBooks

		selectFields := "books.cover_photo_url, books.id, books.title, SUM(order_details.amount) AS total_sold, order_details.amount, books.rating"
		joinSentence := "JOIN order_details ON order_details.book_id = books.id"

		err := database.DB.Table("books").
			Joins(joinSentence).
			Select(selectFields).
			Order("total_sold DESC").
			Group("books.id, order_details.amount").
			Scan(&mostSelledBooks)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code" : consts.ErrorCodeDatabase,
				"details": err.Error.Error(),
				"target" : consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": mostSelledBooks,
		})

	}
}

type Formats struct {
	ID uint `json:"id"`
	Price  float64 `json:"price"`

}

func (controller *BookController) GetBookFormats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookWithFormat usermodels.Book

		id := ctx.Params.ByName("id")


		err := database.DB.Model(&usermodels.Book{}).Joins("AudioFormat").Joins("HardCoverFormat").Joins("DigitalFormat").Where("books.id = ?", id).Find(&bookWithFormat)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code" : consts.ErrorCodeDatabase,
				"error": err.Error.Error(),
				"target" : consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"formats": bookWithFormat,
		})


	}
}
