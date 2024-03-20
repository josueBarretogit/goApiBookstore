package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	bookmodels "api/bookstoreApi/models/bookModels"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AudioFormatController struct {
	GenericController[bookmodels.AudioBookFormat]
}

type AudioDetailsDTO struct {
	Duration    string `json:"duration"`
	ProgramType string `json:"program_type"`
}

func (controller *AudioFormatController) GetAudioDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var details AudioDetailsDTO

		if !helpers.IsNumber(id) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.AudioBookFormatModelName,
			})
			return
		}

		sqlSentence := helpers.NewSQLBuilder("audio_book_formats").
			Select("duration", "program_type").
			Where("audio_book_formats.id = $1").
			AndWhere("audio_book_formats.deleted_at IS NULL").
			GetSQL()

		err := database.Pg.QueryRow(context.Background(), sqlSentence, id).Scan(&details.Duration, &details.ProgramType)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.AudioBookFormatModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"details": details,
		})
	}
}
