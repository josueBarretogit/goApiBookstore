package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"
	"time"

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
				"code":    consts.ErrorCodeDatabase,
				"details": err.Error.Error(),
				"target":  consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": mostSelledBooks,
		})
	}
}

type FormatDTO struct {
	ID    uint    `json:"id"`
	Price float64 `json:"price"`
}

func (controller *BookController) GetBookFormats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var digitalFormat FormatDTO
		var audioFormat FormatDTO
		var hardCover FormatDTO

		id := ctx.Params.ByName("id")
		err := database.DB.Table("books").
			Select("audio_book_formats.price, audio_book_formats.id").
			Joins("LEFT JOIN audio_book_formats ON audio_book_formats.book_id = books.id").
			Where("books.id = ?", id).
			Scan(&audioFormat)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		err = database.DB.Table("books").
			Select("digital_formats.price, digital_formats.id").
			Joins("LEFT JOIN digital_formats ON digital_formats.book_id = books.id").
			Where("books.id = ?", id).
			Scan(&digitalFormat)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		err = database.DB.Table("books").
			Select("hard_cover_formats.price, hard_cover_formats.id").
			Joins("LEFT JOIN hard_cover_formats ON hard_cover_formats.book_id = books.id").
			Where("books.id = ?", id).
			Scan(&hardCover)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"digital":   digitalFormat,
			"audio":     audioFormat,
			"hardCover": hardCover,
		})
	}
}

type GetReviewDto struct {
	Rating     int    `json:"rating,omitempty"`
	Title      string `json:"title,omitempty"`
	BodyReview string `json:"body_review,omitempty"`
	Customer   struct {
		ProfilePictureUrl string `json:"profile_picture_url"`
		Account           struct {
			Username string `json:"username"`
		} `json:"account" gorm:"embedded"`
	} `json:"customer" gorm:"embedded"`
}

func (controller *BookController) GetReviews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reviews []GetReviewDto

		sl := "reviews.rating as rating, reviews.title as title, reviews.body_review as body_review, accounts.username as username, customers.profile_picture_url as profile_picture_url"

		id := ctx.Params.ByName("id")
		err := database.DB.Table("books").
			Select(sl).
			Joins("INNER JOIN reviews on reviews.book_id = books.id").
			Joins("INNER JOIN customers on reviews.customer_id = customers.id").
			Joins("INNER JOIN accounts on accounts.id = customers.account_id").
			Where("books.id = ?", id).
			Scan(&reviews)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"reviews": reviews,
		})
	}
}

type Authors struct {
	ID       uint
	Name     string `json:"name" `
	Lastname string `json:"lastname"`
}

type SearchBookDTO struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Rating          int       `json:"rating"`
	CoverPhotoUrl   string    `json:"cover_photo_url"`
	PublicationDate time.Time `json:"publication_date,omitempty"`
}

type SearchedBook struct {
	Book SearchBookDTO
	Authors []Authors
}

func (controller *BookController) SearchBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var books []SearchBookDTO
		var searchedBooks [5]SearchedBook

		err := database.DB.Table("books").
			Select("books.id as id, books.title as title, books.rating as rating, books.cover_photo_url as cover_photo_url, books.publication_date as publication_date").
			Order("books.id desc").
			Limit(5).
			Scan(&books)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"target": controller.ModelName,
				"error":  err.Error.Error(),
			})
			return
		}


		for index, book := range books {
			var authorsAssociated []Authors
			err = database.DB.Table("authors").
				Select("authors.name as name, authors.id as id").
				Joins("INNER JOIN author_book ON authors.id = author_book.author_id ").
				Joins("INNER JOIN books ON books.id = author_book.book_id ").
				Where("books.id = ?" , book.ID).
				Find(&authorsAssociated)

			if err.Error != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   consts.ErrorCodeDatabase,
					"target": controller.ModelName,
					"error":  err.Error.Error(),
				})
				return
			}
			searchedBooks[index].Authors = authorsAssociated
			searchedBooks[index].Book.ID = book.ID
			searchedBooks[index].Book.Title= book.Title
			searchedBooks[index].Book.Rating = book.Rating
			
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": searchedBooks,
			
		})
	}
}
