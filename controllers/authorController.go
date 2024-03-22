package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	GenericController[usermodels.Author]
}

func (controller *AuthorController) AssignPublisher() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher](controller.RelationName, consts.AuthorModelName)
}

func (controller *AuthorController) AssignBook() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher]("Books", consts.AuthorModelName)
}

type AuthorBooksDTO struct {
	ID            uint       `json:"ID"`
	CoverPhotoUrl string     `json:"coverPhotoUrl"`
	Title         string     `json:"title"`
	Rating        string     `json:"rating"`
	Formats       FormatsDTO `json:"formats"`
}

func (controller *AuthorController) GetAuthorBooks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var books []AuthorBooksDTO

		idAuthor := ctx.Param("id")

		if !helpers.IsNumber(idAuthor) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.BookModelName,
			})
			return
		}

		db, err := database.DB.DB()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.BookModelName,
			})
			return
		}

		sqlSentence := helpers.NewSQLBuilder("books").
			Select(
				`books.id as id`,
				`books.cover_photo_url as cover_photo_url`,
				`books.title as title`,
				consts.AVGrating,
				`hard_cover_formats.price as hard_cover_price`,
				`audio_book_formats.price as audio_book_price`,
				`digital_formats.price as digital_price`,
			).
			InnerJoins("author_book", "author_book.book_id = books.id").
			InnerJoins("authors", "authors.id = author_book.author_id").
			LeftJoins("hard_cover_formats", "hard_cover_formats.book_id = books.id").
			LeftJoins("audio_book_formats", "audio_book_formats.book_id = books.id").
			LeftJoins("digital_formats", "digital_formats.book_id = books.id").
			Where("authors.id = $1").GetSQL()

		rows, err := db.Query(sqlSentence, idAuthor)
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

			var bookFound AuthorBooksDTO
			err = rows.Scan(
				&bookFound.ID,
				&bookFound.CoverPhotoUrl,
				&bookFound.Title,
				&bookFound.Rating,
				&bookFound.Formats.HardCover.Price,
				&bookFound.Formats.AudioFormat.Price,
				&bookFound.Formats.DigitalFormat.Price,
			)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":   consts.ErrorCodeDatabase,
					"error":  err.Error(),
					"target": consts.BookModelName,
				})
				return
			}

			books = append(books, bookFound)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	}
}
