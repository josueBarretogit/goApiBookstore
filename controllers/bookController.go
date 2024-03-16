package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	usermodels "api/bookstoreApi/models/userModels"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	GenericController[usermodels.Book]
}

func (controller *BookController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Book, usermodels.Author]("Authors", consts.BookModelName)
}

func (controller *BookController) AssignLanguage() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Book, usermodels.Language]("Language", consts.BookModelName)
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

type FormatsDTO struct {
	DigitalFormat FormatDTO `gorm:"embedded"`
	AudioFormat FormatDTO `gorm:"embedded"`
	HardCover FormatDTO `gorm:"embedded"`
}

func (controller *BookController) GetBookFormats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var formats FormatsDTO

		id := ctx.Params.ByName("id")
		err := database.DB.Table("books").
			Select("audio_book_formats.price, audio_book_formats.id, digital_formats.price, digital_formats.id,hard_cover_formats.price, hard_cover_formats.id").
			Joins("LEFT JOIN audio_book_formats ON audio_book_formats.book_id = books.id").
			Joins("LEFT JOIN digital_formats ON digital_formats.book_id = books.id").
			Joins("LEFT JOIN hard_cover_formats ON hard_cover_formats.book_id = books.id").
			Where("books.id = ?", id).
			Scan(&formats)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}


		ctx.JSON(http.StatusOK, gin.H{
			"digital":   formats.DigitalFormat,
			"audio":     formats.AudioFormat,
			"hardCover": formats.HardCover,
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
	ID       *string `json:"id"`
	Name     *string `json:"name" `
	Lastname *string `json:"lastname"`
}

type BookDTO struct {
	ID              int          `json:"id"`
	Title           string       `json:"title"`
	Rating          int          `json:"rating"`
	CoverPhotoUrl   string       `json:"cover_photo_url"`
	PublicationDate time.Time    `json:"publication_date"`
	HardCoverPrice  *float64     `json:"hardCoverPrice"`
	DigitalPrice    *float64     `json:"digitalPrice"`
	AudioPrice      *float64     `json:"audioPrice"`
	Authors         []AuthorsDTO `json:"authors"`
}

type BookFilter struct {
	Genre          string
	SearchTerm     string
	RangePrices    string
	RangePrice1    string
	RangePrice2    string
	HighToLowPrice string
	Rating         string
	Languages      string
	DatesFrom      string
	DatesTo        string
	Page           int
	ItemsPerPage   int
}

func BuildSearchBookSql(filters BookFilter) string {
	sqlBuilder := helpers.NewSQLBuilder("books")

	sqlBuild := sqlBuilder.
		Select(
			`COUNT(books.id) OVER() as total`,
			`books.id as id`,
			`books.title as title`,
			`books.rating as rating`,
			`books.cover_photo_url as cover_photo_url`,
			`books.publication_date as publication_date`,
			`ARRAY_AGG(authors.id) as author_ids`,
			`ARRAY_AGG(authors.name) as author_names`,
			`ARRAY_AGG(authors.lastname) as author_lastnames`,
			`hard_cover_formats.price as hard_cover_price`,
			`audio_book_formats.price as audio_book_price`,
			`digital_formats.price as digital_price`).
		LeftJoins(`author_book `, `author_book.book_id = books.id`).
		LeftJoins(`genres`, `genres.id = books.genre_id`).
		LeftJoins(`hard_cover_formats `, `hard_cover_formats.book_id = books.id `).
		LeftJoins(`audio_book_formats`, `audio_book_formats.book_id = books.id `).
		LeftJoins(`digital_formats`, `digital_formats.book_id = books.id `).
		LeftJoins(`authors`, `author_book.author_id =  authors.id  `).
		Where(`(authors.name LIKE '%' || $1 || '%' OR books.title LIKE '%' || $1 || '%' )`)

	if filters.Genre != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`books.genre_id = %s`, filters.Genre))
	}

	if filters.RangePrice1 != "undefined" && filters.RangePrice2 != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`hard_cover_formats.price BETWEEN %s AND %s`, filters.RangePrice1, filters.RangePrice2))
	}

	if filters.Rating != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`  books.rating >= %s`, filters.Rating))
	}

	if filters.DatesFrom != "undefined" && filters.DatesTo != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(` books.publication_date BETWEEN '%s' AND '%s'`, filters.DatesFrom, filters.DatesTo))
	} else if filters.DatesFrom != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`books.publication_date >= '%s'`, filters.DatesFrom))
	} else if filters.DatesTo != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(` AND books.publication_date <= '%s'`, filters.DatesTo))
	}

	sqlBuild.Group().
		BY("books.id, ").
		BY("hard_cover_formats.price,").
		BY("audio_book_formats.price,").
		BY("digital_formats.price")

	if filters.RangePrices == "yes" {
		sqlBuild.OrderBy(` hard_cover_formats.price DESC, digital_formats.price DESC, audio_book_formats.price DESC`)
	} else {
		sqlBuild.OrderBy(` hard_cover_formats.price ASC, digital_formats.price ASC, audio_book_formats.price ASC`)
	}

	sqlBuild.
		Paginate(filters.Page, filters.ItemsPerPage)

	return sqlBuild.GetSQL()
}

func (controller *BookController) SearchBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Param("page")
		itemsPerPage := ctx.Param("itemsPerPage")
		searchTerm := ctx.Param("searchTerm")
		genreId := ctx.Param("genre")
		rangePrice1 := ctx.Param("rangePrice1")
		rangePrice2 := ctx.Param("rangePrice2")
		highToLowPrice := ctx.Param("highToLowPrice")
		rating := ctx.Param("rating")
		fromDate := ctx.Param("fromDate")
		ToDate := ctx.Param("toDate")

		var booksPg []BookDTO

		pageInt, errConv := strconv.Atoi(page)
		if errConv != nil {
			return
		}
		itemsPerPageInt, errConv := strconv.Atoi(itemsPerPage)
		if errConv != nil {
			return
		}

		if !helpers.IsNumberOrUndefined(genreId) || !helpers.IsNumberOrUndefined(rangePrice1) || !helpers.IsNumberOrUndefined(rangePrice2) ||
			!helpers.IsNumberOrUndefined(rating) || !helpers.IsNumberOrUndefined(page) || !helpers.IsNumberOrUndefined(itemsPerPage) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		if !helpers.IsDateOrUndefined(fromDate) || !helpers.IsDateOrUndefined(ToDate) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorBadDate,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		filters := BookFilter{
			SearchTerm:   searchTerm,
			Genre:        genreId,
			RangePrices:  highToLowPrice,
			RangePrice1:  rangePrice1,
			RangePrice2:  rangePrice2,
			Rating:       rating,
			Languages:    ``,
			DatesFrom:    fromDate,
			DatesTo:      ToDate,
			Page:         pageInt,
			ItemsPerPage: itemsPerPageInt,
		}

		if filters.SearchTerm == "undefined" {
			filters.SearchTerm = ""
		}

		sqlSentence := BuildSearchBookSql(filters)

		rows, err := database.Pg.Query(context.Background(), sqlSentence, filters.SearchTerm)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return

		}
		defer rows.Close()

		var total int
		for rows.Next() {
			var book BookDTO
			var authorIDs []*string
			var authorNames []*string
			var authorLastNames []*string

			err = rows.Scan(
				&total,
				&book.ID,
				&book.Title,
				&book.Rating,
				&book.CoverPhotoUrl,
				&book.PublicationDate,
				&authorIDs,
				&authorNames,
				&authorLastNames,
				&book.HardCoverPrice,
				&book.AudioPrice,
				&book.DigitalPrice,
			)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   consts.ErrorCodeDatabase,
					"error":  err.Error(),
					"target": consts.BookModelName,
				})
				return
			}

			for i, id := range authorIDs {
				if id != nil {
					book.Authors = append(book.Authors, AuthorsDTO{
						ID:       id,
						Name:     authorNames[i],
						Lastname: authorLastNames[i],
					})
				} else {
					book.Authors = []AuthorsDTO{}
				}
			}
			booksPg = append(booksPg, book)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books":      booksPg,
			"totalBooks": total,
		})
	}
}
