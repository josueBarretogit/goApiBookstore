package controllers

import (
	"net/http"

	"api/bookstoreApi/database"
	"api/bookstoreApi/helpers"
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomerController struct {
	GenericController[usermodels.Customer]
}

type AuthorController struct {
	GenericController[usermodels.Author]
}

type AccountController struct {
	GenericController[usermodels.Account]
}

type RoleController struct {
	GenericController[usermodels.Role]
}

type PublisherController struct {
	GenericController[usermodels.Publisher]
}

type BookController struct {
	GenericController[bookmodels.Book]
}

type HardCoverFormatController struct {
	GenericController[bookmodels.HardCoverFormat]
}

type DigitalFormatController struct {
	GenericController[bookmodels.DigitalFormat]
}

type AudioBookFormatController struct {
	GenericController[bookmodels.AudioBookFormat]
}
type OrderController struct {
	GenericController[paymentmodels.Order]
}

type PaymentController struct {
	GenericController[paymentmodels.Payment]
}

type PurchaseDetailsController struct {
	GenericController[paymentmodels.OrderDetails]
}

type CreditCardController struct {
	GenericController[paymentmodels.CreditCard]
}

type BankAccountController struct {
	GenericController[paymentmodels.BankAccount]
}

type ReviewController struct {
	GenericController[paymentmodels.Review]
}

func (controller *PublisherController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Publisher, usermodels.Author](controller.RelationName)
}

func (controller *AccountController) AssignRole() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Account, usermodels.Role](controller.RelationName)
}

func (controller *AuthorController) AssignPublisher() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher](controller.RelationName)
}

func (controller *AccountController) LogIn() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		newToken, err  := helpers.GenerateNewJwtToken(jwt.MapClaims{
			"accountId" : "a",
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}


		ctx.JSON(http.StatusOK, gin.H{
			"token" : newToken,
		})

	}
}

func AssignManyToManyRelation[T interface{}, K interface{}](relation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelToUpdate T
		var modelData K

		id := c.Params.ByName("id")
		err := database.DB.First(&modelToUpdate, id)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		c.BindJSON(&modelData)
		errDatabase := database.DB.Model(&modelToUpdate).Association(relation).Append(&modelData)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": errDatabase.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"updated": modelData,
		})
		return
	}
}
