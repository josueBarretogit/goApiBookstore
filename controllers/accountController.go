package controllers

import (
	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (controller *AccountController) LogIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var payload usermodels.Account
		var accountFound usermodels.Account

		errJson := ctx.BindJSON(&payload)


		if errJson != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error" : errJson.Error(),
			})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		errDb := database.DB.Model(&payload).Where(&payload.Username).Find(&accountFound)

		if errDb.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error" : errDb.Error.Error(),
			})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		verifiyPassowrd := helpers.CheckPasswordHash(accountFound.Password, payload.Password)
		
		if !verifiyPassowrd {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response" : "This account doesnt exist",
				"success" :false,
			})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		newToken, err  := helpers.GenerateNewJwtToken(jwt.MapClaims{
			"accountId" : payload.ID,
			"username" : payload.Username,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response" : "login succesfull",
			"success" : true,
			"token" : newToken,
			"accountLogged":  payload,
		})

	}
}

func (controller *AccountController) AssignRole() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Account, usermodels.Role](controller.RelationName)
}

func (controller *AccountController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var newAccount usermodels.Account

		errPayload := ctx.BindJSON(&newAccount)

		if errPayload != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Received bad data",
				"details": errPayload.Error(),
			})
			return
		}
		
		HashedPassowrd , errHashing := helpers.HashPassword(newAccount.Password)
		

		if errHashing != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": errHashing.Error(),
			})
		}

		newAccount.Password = HashedPassowrd

		err := database.DB.Create(&newAccount)
		if err.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"registered": newAccount,
		})
		return

	}
}



