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

type BestSellerBooks struct {
	ID            uint   `json:"ID"`
	Title         string `json:"title"`
	CoverPhotoUrl string `json:"coverPhotoUrl"`
	Rating        *int   `json:"rating,omitempty"`
	TotalSold     int    `json:"totalSold"`
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

		selectFields := `books.cover_photo_url, books.id, books.title,(SELECT SUM(sold) from UNNEST(ARRAY_AGG(CAST(order_details.amount as INT))) sold) as total_sold,  books.rating`
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
	DigitalFormat FormatDTO `gorm:"embedded"`
	AudioFormat   FormatDTO `gorm:"embedded"`
	HardCover     FormatDTO `gorm:"embedded"`
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
		ProfilePictureUrl string `json:"profilePictureUrl"`
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
		LeftJoins(`authors`, `author_book.author_id =  authors.id  `).
		LeftJoins(`genres`, `genres.id = books.genre_id`).
		LeftJoins(`hard_cover_formats `, `hard_cover_formats.book_id = books.id `).
		LeftJoins(`audio_book_formats`, `audio_book_formats.book_id = books.id `).
		LeftJoins(`digital_formats`, `digital_formats.book_id = books.id `).
		LeftJoins(`language_book`, `language_book.book_id =  books.id  `).
		LeftJoins(`languages`, `language_book.language_id =  languages.id  `).
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
		sqlBuild.AndWhere(fmt.Sprintf(`  languages.id IN (%s)`, filters.Languages))
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
		BY("languages.id, ").
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
	Rating          int           `json:"rating"`
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
		)
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
	a := 3

	for i := 0; i < 2000; i++ {
		database.DB.Model(&usermodels.Book{}).Create(&usermodels.Book{
			Title:         faker.Name(),
			CoverPhotoUrl: faker.URL(),
			Description:   faker.Paragraph(),
			Rating:        &a,
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

// BACKUP SENA

// SELECT books.cover_photo_url, books.id, books.title,
// (SELECT SUM(sold) from UNNEST(ARRAY_AGG(CAST(order_details.amount as int))) sold)
// as total_sold, books.rating FROM books INNER JOIN order_details ON order_details.book_id = books.id GROUP BY books.id ORDER BY id desc
