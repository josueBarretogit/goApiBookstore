package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"context"
	"fmt"
	"net/http"
	"strings"
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

type AuthorsDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name" `
	Lastname string `json:"lastname"`
}

type BookDTO struct {
	ID              int          `json:"id"`
	Title           string       `json:"title"`
	Rating          int          `json:"rating"`
	CoverPhotoUrl   string       `json:"cover_photo_url"`
	PublicationDate time.Time    `json:"publication_date,omitempty"`
	Authors         []AuthorsDTO `json:"authors"`
}

func (controller *BookController) SearchBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchTerm := strings.Replace(ctx.Param("searchTerm"), "/", "", -1)
		fmt.Println(searchTerm)
		var booksPg []BookDTO

		selectSentence := `
            SELECT 
                books.id as id, 
                books.title as title, 
                books.rating as rating , 
                books.cover_photo_url as cover_photo_url, 
                books.publication_date as publication_date,
                ARRAY_AGG(authors.id) as author_ids,
                ARRAY_AGG(authors.name) as author_names,
                ARRAY_AGG(authors.lastname) as author_lastnames
            FROM 
                books
            INNER JOIN 
                author_book ON author_book.book_id = books.id
            INNER JOIN 
                authors ON author_book.author_id = authors.id`
		selectSentence += ` WHERE authors.name LIKE ` + `'%` + searchTerm + `%'` + ` GROUP BY books.id`

		rows, err := database.Pg.Query(context.Background(), selectSentence)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var book BookDTO
			var authorIDs []int
			var authorNames []string
			var authorLastNames []string

			err = rows.Scan(&book.ID, &book.Title, &book.Rating, &book.CoverPhotoUrl, &book.PublicationDate, &authorIDs, &authorNames, &authorLastNames)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   consts.ErrorCodeDatabase,
					"error":  err.Error(),
					"target": consts.BookModelName,
				})
				return
			}

			for i, id := range authorIDs {
				book.Authors = append(book.Authors, AuthorsDTO{
					ID:       id,
					Name:     authorNames[i],
					Lastname: authorLastNames[i],
				})
			}

			booksPg = append(booksPg, book)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": booksPg,
		})
	}
}
