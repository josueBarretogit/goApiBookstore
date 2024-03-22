package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	GenericController[usermodels.Customer]
}

type BookReviewCustomerDTO struct {
	BookID        uint   `json:"ID"`
	CoverPhotoUrl string `json:"coverPhotoUrl"`
}

type CustomerReviewDTO struct {
	ID     uint                  `json:"ID"`
	Rating string                `json:"rating"`
	Title  string                `json:"title"`
	Body   string                `json:"body"`
	Book   BookReviewCustomerDTO `json:"book" gorm:"embedded"`
}

func (controller *CustomerController) GetCustomerReviews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reviews []CustomerReviewDTO
		idCustomer := ctx.Param("id")

		if !helpers.IsNumber(idCustomer) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		err := database.DB.Table("customers").
			Select(
				`customers.id as id,
				reviews.rating as rating, 
				reviews.title as title,
				reviews.body_review as body,
				books.id as book_id,
				books.cover_photo_url as cover_photo_url
			`,
			).
			Joins("INNER JOIN reviews ON reviews.customer_id = customers.id").
			Joins("INNER JOIN books ON books.id = reviews.book_id").
			Where("customers.id = ? ", idCustomer).
			Find(&reviews)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.CustomerModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"reviews": reviews,
		})
	}
}
