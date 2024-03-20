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

type HardCoverFormatController struct {
	GenericController[bookmodels.HardCoverFormat]
}

type HardCoverFormatDetailsDto struct {
	Height   float32 `json:"height,"`
	Width    float32 `json:"width,"`
	Length   float32 `json:"length,"`
	NumPages int     `json:"numPages,"`
	Stock    int64   `json:"stock,"`
	Weight   float64 `json:"weight,"`
}

func (controller *HardCoverFormatController) GetHardCoverFormatDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var details HardCoverFormatDetailsDto

		if !helpers.IsNumber(id) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":   consts.ErrorNotNumber,
				"error":  "",
				"target": consts.HardcoverFormatModelName,
			})
			return
		}

		sqlSentence := helpers.NewSQLBuilder("hard_cover_formats").
			Select("height",
				"width",
				"length",
				"num_pages",
				"stock",
				"weight",
			).
			Where("hard_cover_formats.id = $1").
			AndWhere("hard_cover_formats.deleted_at IS NULL").
			GetSQL()

		err := database.Pg.QueryRow(context.Background(), sqlSentence, id).Scan(
			&details.Height,
			&details.Width,
			&details.Length,
			&details.NumPages,
			&details.Stock,
			&details.Weight,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   consts.ErrorCodeDatabase,
				"error":  err.Error(),
				"target": consts.HardcoverFormatModelName,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"details": details,
		})
	}
}
