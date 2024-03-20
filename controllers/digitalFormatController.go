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

type DigitalFormatController struct {
	GenericController[bookmodels.DigitalFormat]
}

type DigitalFormatDto struct {
	ScreenReader bool `json:"screenReader"`
	TextToSpeech bool `json:"textToSpeech"`
}

func (controller *DigitalFormatController) GetDigitalFormatDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var details DigitalFormatDto

		if !helpers.IsNumber(id) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.AudioBookFormatModelName,
			})
			return
		}

		sqlSentence := helpers.NewSQLBuilder("digital_formats").
			Select("screen_reader", "text_to_speech").
			Where("digital_formats.id = $1").
			AndWhere("digital_formats.deleted_at IS NULL").
			GetSQL()

		err := database.Pg.QueryRow(context.Background(), sqlSentence, id).Scan(&details.ScreenReader, &details.TextToSpeech)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.DigitalFormatModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"details": details,
		})
	}
}
