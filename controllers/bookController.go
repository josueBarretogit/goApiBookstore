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
	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5"
)

type BookController struct {
	GenericController[usermodels.Book]
}

func (controller *BookController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Book, usermodels.Author]("Authors", consts.BookModelName)
}

func (controller *BookController) AssignLanguage() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Book, usermodels.Language]("Languages", consts.BookModelName)
}

type AuthorsDTO struct {
	ID       *string `json:"id"`
	Name     *string `json:"name" `
	Lastname *string `json:"lastname"`
}

type BookDTO struct {
	ID              int          `json:"id"`
	Title           string       `json:"title"`
	Rating          *string      `json:"rating"`
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

type BestSellerBooks struct {
	ID            uint    `json:"ID"`
	Title         string  `json:"title"`
	CoverPhotoUrl string  `json:"coverPhotoUrl"`
	Rating        *string `json:"rating,omitempty"`
	TotalSold     int     `json:"totalSold"`
}

func (controller *BookController) GetBestSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mostSelledBooks []BestSellerBooks
		itemsPerPage := ctx.Param("itemsPerPage")
		page := ctx.Param("page")
		idGenre := ctx.Param("idGenre")

		if !helpers.IsNumber(page) || !helpers.IsNumber(itemsPerPage) || !helpers.IsNumberOrUndefined(idGenre) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		pageInt, errConv := strconv.Atoi(page)
		if errConv != nil {
			return
		}
		itemsPerPageInt, errConv := strconv.Atoi(itemsPerPage)
		if errConv != nil {
			return
		}

		selectFields := `
		books.cover_photo_url,
		books.id, books.title,
		(SELECT SUM(sold) from UNNEST(ARRAY_AGG(CAST(order_details.amount as INT))) sold) as total_sold,
		` + consts.AVGrating

		joinSentence := "INNER JOIN order_details ON order_details.book_id = books.id"

		queryBuilder := database.DB.Table("books").
			Select(selectFields).
			Joins(joinSentence).
			Joins("INNER JOIN genres ON books.genre_id = genres.id").
			Order("total_sold DESC").
			Offset((pageInt - 1) * itemsPerPageInt).
			Limit(itemsPerPageInt)

		if idGenre != "undefined" {
			queryBuilder.Where("genres.id = ?", idGenre)
		}

		queryBuilder.Group("books.id, genres.id")

		err := queryBuilder.Scan(&mostSelledBooks)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"details": err.Error.Error(),
				"target":  consts.BookModelName,
			})
			return
		}

		if mostSelledBooks == nil {
			mostSelledBooks = []BestSellerBooks{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": mostSelledBooks,
		})
	}
}

type FormatDTO struct {
	ID    uint    `json:"ID"`
	Price float64 `json:"price"`
}

type FormatsDTO struct {
	DigitalFormat FormatDTO `json:"digitalFormat" gorm:"embedded"`
	AudioFormat   FormatDTO `json:"audioFormat" gorm:"embedded"`
	HardCover     FormatDTO `json:"hardCover" gorm:"embedded"`
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
	BodyReview string `json:"bodyReview,omitempty"`
	Customer   struct {
		ID                uint   `json:"ID"`
		ProfilePictureUrl string `json:"profilePictureUrl"`
		Account           struct {
			Username string `json:"username"`
		} `json:"account" gorm:"embedded"`
	} `json:"customer" gorm:"embedded"`
}

func (controller *BookController) GetReviews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reviews []GetReviewDto

		sl := `reviews.rating as rating, 
		reviews.title as title,
		reviews.body_review as body_review,
		accounts.username as username,
		customers.profile_picture_url as profile_picture_url,
		customers.id as id
		`

		id := ctx.Params.ByName("id")
		page := ctx.Params.ByName("page")
		itemsPerPages := ctx.Params.ByName("itemsPerPages")
		rating := ctx.Params.ByName("rating")

		if !helpers.IsNumber(id) || !helpers.IsNumber(page) || !helpers.IsNumber(itemsPerPages) || !helpers.IsNumberOrUndefined(rating) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.HardcoverFormatModelName,
			})
			return
		}

		pageInt, errConv := strconv.Atoi(page)
		if errConv != nil {
			return
		}
		itemsPerPageInt, errConv := strconv.Atoi(itemsPerPages)
		if errConv != nil {
			return
		}

		sqlBuilder := database.DB.Table("books").
			Select(sl).
			Joins("INNER JOIN reviews on reviews.book_id = books.id").
			Joins("INNER JOIN customers on reviews.customer_id = customers.id").
			Joins("INNER JOIN accounts on accounts.id = customers.account_id").
			Offset((pageInt-1)*itemsPerPageInt).
			Limit(itemsPerPageInt).
			Where("books.id = ? ", id)

		if rating != "undefined" {
			sqlBuilder.
				Where("reviews.rating = ?", rating)
		}
		sqlBuilder.Where("books.deleted_at IS NULL")

		err := sqlBuilder.Scan(&reviews)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		if reviews == nil {
			reviews = []GetReviewDto{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"reviews": reviews,
		})
	}
}

type BookReviewDTO struct {
	ID            uint         `json:"ID"`
	Title         string       `json:"title"`
	CoverPhotoUrl string       `json:"coverPhotoUrl"`
	Authors       []AuthorsDTO `json:"authors"`
}

type RatingsStatisticsDTO struct {
	FiveStar struct {
		Percentage *float64 `json:"percentage"`
	} `json:"fiveStar"`
	FourStar struct {
		Percentage *float64 `json:"percentage"`
	} `json:"fourStar"`
	ThreeStar struct {
		Percentage *float64 `json:"percentage"`
	} `json:"threeStar"`
	TwoStart struct {
		Percentage *float64 `json:"percentage"`
	} `json:"twoStar"`
	OneStar struct {
		Percentage *float64 `json:"percentage"`
	} `json:"oneStar"`
}

type ReviewStatisticsDTO struct {
	Book         BookReviewDTO        `json:"book"`
	Ratings      RatingsStatisticsDTO `json:"ratings"`
	TotalReviews *float64             `json:"totalReviews"`
}

func GetReviewStatisticsSql() string {
	sqlBuilder := helpers.NewSQLBuilder("books").
		Select(
			"books.id as id",
			"books.title as title",
			"books.cover_photo_url as cover_photo_url",
			"ARRAY_AGG(DISTINCT authors.id) as author_id",
			"ARRAY_AGG(DISTINCT authors.name) as name",
			"ARRAY_AGG(DISTINCT authors.lastname) as lastname",
			"COUNT(DISTINCT reviews.id) as total_rating",
			`(Select COUNT(reviews.rating) FROM reviews INNER JOIN books ON reviews.book_id = books.id WHERE books.id = $1 AND reviews.rating = 5 AND books.deleted_at IS NULL) as total_five_rating`,
			"(Select COUNT(reviews.rating) FROM reviews INNER JOIN books ON reviews.book_id = books.id WHERE books.id = $1 AND reviews.rating = 4 AND books.deleted_at IS NULL) as total_four_rating",
			"(Select COUNT(reviews.rating) FROM reviews INNER JOIN books ON reviews.book_id = books.id WHERE books.id = $1 AND reviews.rating = 3 AND books.deleted_at IS NULL) as total_three_rating",
			"(Select COUNT(reviews.rating) FROM reviews INNER JOIN books ON reviews.book_id = books.id WHERE books.id = $1 AND reviews.rating = 2 AND books.deleted_at IS NULL) as total_two_rating",
			"(Select COUNT(reviews.rating) FROM reviews INNER JOIN books ON reviews.book_id = books.id WHERE books.id = $1 AND reviews.rating = 1 AND books.deleted_at IS NULL)  as total_one_rating",
		).
		InnerJoins("reviews", "reviews.book_id = books.id").
		InnerJoins("author_book", "author_book.book_id = books.id").
		InnerJoins("authors", "authors.id = author_book.author_id").
		Where("books.id = $1").
		AndWhere("books.deleted_at IS NULL").
		Group().BY("books.id")

	return sqlBuilder.GetSQL()
}

func (controller *BookController) GetReviewStatistics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idBook := ctx.Param("idBook")

		if !helpers.IsNumber(idBook) {

			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.HardcoverFormatModelName,
			})
			return
		}

		var ReviewStatistics ReviewStatisticsDTO
		var authorIDs []*string
		var authorsName []*string
		var authorsLastName []*string

		sqlSentence := GetReviewStatisticsSql()

		err := database.Pg.QueryRow(context.Background(), sqlSentence, idBook).Scan(
			&ReviewStatistics.Book.ID,
			&ReviewStatistics.Book.Title,
			&ReviewStatistics.Book.CoverPhotoUrl,
			&authorIDs,
			&authorsName,
			&authorsLastName,
			&ReviewStatistics.TotalReviews,
			&ReviewStatistics.Ratings.FiveStar.Percentage,
			&ReviewStatistics.Ratings.FourStar.Percentage,
			&ReviewStatistics.Ratings.ThreeStar.Percentage,
			&ReviewStatistics.Ratings.TwoStart.Percentage,
			&ReviewStatistics.Ratings.OneStar.Percentage,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		for index, id := range authorIDs {
			ReviewStatistics.Book.Authors = append(ReviewStatistics.Book.Authors, AuthorsDTO{
				ID:       id,
				Name:     authorsName[index],
				Lastname: authorsLastName[index],
			})
		}

		if ReviewStatistics.Ratings.OneStar.Percentage != nil {
			result1 := (*ReviewStatistics.Ratings.OneStar.Percentage / *ReviewStatistics.TotalReviews) * 100
			ReviewStatistics.Ratings.OneStar.Percentage = &result1
		}

		if ReviewStatistics.Ratings.TwoStart.Percentage != nil {
			result1 := (*ReviewStatistics.Ratings.TwoStart.Percentage / *ReviewStatistics.TotalReviews) * 100
			ReviewStatistics.Ratings.TwoStart.Percentage = &result1
		}

		if ReviewStatistics.Ratings.ThreeStar.Percentage != nil {
			result1 := (*ReviewStatistics.Ratings.ThreeStar.Percentage / *ReviewStatistics.TotalReviews) * 100
			ReviewStatistics.Ratings.ThreeStar.Percentage = &result1
		}

		if ReviewStatistics.Ratings.FourStar.Percentage != nil {
			result1 := (*ReviewStatistics.Ratings.FourStar.Percentage / *ReviewStatistics.TotalReviews) * 100
			ReviewStatistics.Ratings.FourStar.Percentage = &result1
		}

		if ReviewStatistics.Ratings.FiveStar.Percentage != nil {
			result1 := (*ReviewStatistics.Ratings.FiveStar.Percentage / *ReviewStatistics.TotalReviews) * 100
			ReviewStatistics.Ratings.FiveStar.Percentage = &result1
		}

		ctx.JSON(http.StatusOK, gin.H{
			"reviewStatistics": ReviewStatistics,
		})
	}
}

func BuildSearchBookSql(filters BookFilter) string {
	sqlBuilder := helpers.NewSQLBuilder("books")

	sqlBuild := sqlBuilder.
		Select(
			`COUNT(books.id) OVER() as total`,
			`books.id as id`,
			`books.title as title`,
			consts.AVGrating,
			`books.cover_photo_url as cover_photo_url`,
			`books.publication_date as publication_date`,
			`ARRAY_AGG(authors.id) as author_ids`,
			`ARRAY_AGG(authors.name) as author_names`,
			`ARRAY_AGG(authors.lastname) as author_lastnames`,
			`hard_cover_formats.price as hard_cover_price`,
			`audio_book_formats.price as audio_book_price`,
			`digital_formats.price as digital_price`).
		LeftJoins(`author_book `, `author_book.book_id = books.id`).
		LeftJoins(`authors`, `author_book.author_id =  authors.id  `).
		LeftJoins(`genres`, `genres.id = books.genre_id`).
		LeftJoins(`hard_cover_formats `, `hard_cover_formats.book_id = books.id `).
		LeftJoins(`audio_book_formats`, `audio_book_formats.book_id = books.id `).
		LeftJoins(`digital_formats`, `digital_formats.book_id = books.id `).
		Where(`(authors.name ILIKE '%' || $1 || '%' OR books.title ILIKE '%' || $1 || '%' )`)

	if filters.Genre != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`books.genre_id = %s`, filters.Genre))
	}

	if filters.RangePrice1 != "undefined" && filters.RangePrice2 != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`hard_cover_formats.price BETWEEN %s AND %s`, filters.RangePrice1, filters.RangePrice2))
	}

	if filters.Rating != "undefined" {
		sqlBuild.AndWhere(fmt.Sprintf(`  books.rating >= %s`, filters.Rating))
	}

	if filters.Languages != "undefined" {
		sqlBuild.
			LeftJoins(`language_book`, `language_book.book_id =  books.id  `).
			LeftJoins(`languages`, `language_book.language_id =  languages.id  `).
			AndWhere(fmt.Sprintf(`  languages.id IN (%s)`, filters.Languages))
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
		Paginate(filters.Page, filters.ItemsPerPage).
		AndWhere("books.deleted_at IS NULL")

	return sqlBuild.GetSQL()
}

// Filter
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
		idLanguage := ctx.Param("idLanguage")

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
			!helpers.IsNumberOrUndefined(rating) || !helpers.IsNumberOrUndefined(page) || !helpers.IsNumberOrUndefined(itemsPerPage) || !helpers.IsNumberOrUndefined(idLanguage) {
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
			Languages:    idLanguage,
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

		if booksPg == nil {
			booksPg = []BookDTO{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books":      booksPg,
			"totalBooks": total,
		})
	}
}

type LanguageDto struct {
	ID   *string `json:"ID"`
	Name *string `json:"name"`
}

type OneBookDTO struct {
	ID              uint          `json:"ID"`
	Title           string        `json:"title"`
	PublicationDate time.Time     `json:"publicationDate"`
	CoverPhotoUrl   string        `json:"coverPhotoUrl"`
	Rating          *string       `json:"rating"`
	Description     string        `json:"description"`
	Isbn            string        `json:"isbn"`
	Authors         []AuthorsDTO  `json:"authors" `
	Languages       []LanguageDto `json:"languages" `
	Genre           struct {
		Name string `json:"name"`
	} `json:"genre"`
}

func buildGetOneBookSql() string {
	sqlBuilder := helpers.NewSQLBuilder("books")

	sqlSentence := sqlBuilder.Select(
		"books.id as id",
		"books.title as title",
		"books.publication_date as publication_date",
		"books.cover_photo_url as cover_photo_url",
		`ARRAY_AGG(DISTINCT ( authors.id )) as author_ids`,
		`ARRAY_AGG(DISTINCT ( authors.name )) as author_names`,
		`ARRAY_AGG(DISTINCT ( authors.lastname )) as author_lastnames`,
		`ARRAY_AGG(DISTINCT ( languages.id )) as languages_id`,
		`ARRAY_AGG(DISTINCT ( languages.name )) as languages_name`,
		`genres.name as genre_name`,
		consts.AVGrating,
	).
		InnerJoins("author_book", "author_book.book_id = books.id").
		InnerJoins("authors", "authors.id = author_book.author_id").
		InnerJoins("language_book", "language_book.book_id = books.id").
		InnerJoins("languages", "languages.id = language_book.language_id").
		InnerJoins("genres", "genres.id = books.genre_id").
		Where("books.id = $1").
		AndWhere("books.deleted_at IS NULL").
		Group().BY("books.id,").BY("genres.id").
		GetSQL()

	return sqlSentence
}

func (controller *BookController) GetOneBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var book OneBookDTO
		var authorsId []*string
		var authorsName []*string
		var authorsLastName []*string
		var languagesId []*string
		var languagesNames []*string

		idBook := ctx.Param("idBook")

		if !helpers.IsNumber(idBook) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		sqlSentence := buildGetOneBookSql()

		err := database.Pg.QueryRow(context.Background(), sqlSentence, idBook).Scan(
			&book.ID,
			&book.Title,
			&book.PublicationDate,
			&book.CoverPhotoUrl,
			&authorsId,
			&authorsName,
			&authorsLastName,
			&languagesId,
			&languagesNames,
			&book.Genre.Name,
			&book.Rating,
		)

		if err != nil && err.Error() == pgx.ErrNoRows.Error() {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":   pgx.ErrNoRows,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		for index, id := range authorsId {
			book.Authors = append(book.Authors, AuthorsDTO{
				ID:       id,
				Name:     authorsLastName[index],
				Lastname: authorsLastName[index],
			})
		}

		for index, id := range languagesId {
			book.Languages = append(book.Languages, LanguageDto{
				ID:   id,
				Name: languagesNames[index],
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"book": book,
		})
	}
}

func Insert() {
	for i := 0; i < 2000; i++ {
		database.DB.Model(&usermodels.Book{}).Create(&usermodels.Book{
			Title:         faker.Name(),
			CoverPhotoUrl: faker.URL(),
			Description:   faker.Paragraph(),
		})
		fmt.Printf("inserted %d", i)
	}
}

type Test struct {
	Id            uint   `json:"id"`
	Title         string `json:"title"`
	CoverPhotoUrl string `json:"cover_photo"`
	Description   string `json:"description"`
	Rating        uint   `json:"rating"`
}

func (controller *BookController) Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var test []Test

		rows, err := database.Pg.Query(context.Background(), "SELECT id, title, cover_photo_url, description, rating from books where id = 1")
		if err != nil {
			return
		}

		defer rows.Close()

		for rows.Next() {

			var id uint
			var title string
			var CoverPhotoUrl string
			var description string
			var rating uint

			err = rows.Scan(&id, &title, &CoverPhotoUrl, &description, &rating)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   consts.ErrorCodeDatabase,
					"error":  err.Error(),
					"target": consts.BookModelName,
				})
				return
			}

			test = append(test, Test{
				Id:            id,
				Title:         title,
				CoverPhotoUrl: CoverPhotoUrl,
				Description:   description,
				Rating:        rating,
			})

		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": test,
		})
	}
}

func (controller *BookController) UpdateRating() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rating := ctx.Param("rating")
		id := ctx.Param("id")

		var response struct {
			ID uint `json:"ID"`
		}
		err := database.DB.Raw(`
			UPDATE books
			SET rating = rating || ?
			WHERE  id = ?
			RETURNING id
		`, rating, id).Scan(&response)

		if err.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}
