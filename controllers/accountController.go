package controllers

import (
	"net/http"

	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


type AccountController struct {
	GenericController[usermodels.Account]
}


func (controller *AccountController) LogIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload usermodels.Account
		var accountFound usermodels.Account

		errJson := ctx.BindJSON(&payload)

		if errJson != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": errJson.Error(),
			})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		errDb := database.DB.Model(&usermodels.Account{}).Where(usermodels.Account{Username: payload.Username}).Find(&accountFound)

		if errDb.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"dbError": errDb.Error.Error(),
			})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if accountFound.ID == 0 {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": "This account doesnt exist",
				"success":  false,
			})
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		verifiyPassowrd := helpers.CheckPasswordHash(payload.Password, accountFound.Password)

		if !verifiyPassowrd {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": "Password doesnt match",
				"success":  false,
			})
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		newToken, errToken := helpers.GenerateNewJwtToken(jwt.MapClaims{
			"accountID": accountFound.ID,
			"username":  accountFound.Username,
		})

		if errToken != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errorToken": errToken.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response": "login successful",
			"success":  true,
			"token":    newToken,
			"accountLogged": map[string]any{
				"username": accountFound.Username,
				"roles":    accountFound.Roles,
			},
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

		HashedPassowrd, errHashing := helpers.HashPassword(newAccount.Password)

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
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"registered": newAccount,
		})
	}
}
