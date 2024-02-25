package controllers

import (
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	FindAll() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	FindOneBy() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type PublisherController struct {
	GenericController[usermodels.Publisher]
}

func NewPublisherController() *PublisherController {
	generiController := NewGenericController[usermodels.Publisher]("Authors")
	return &PublisherController{
		GenericController: *generiController,
	}
}

func (controller *PublisherController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Publisher, usermodels.Author](controller.RelationName)
}

type RoleController struct {
	GenericController[usermodels.Role]
}

func NewRoleController() *RoleController {
	generiController := NewGenericController[usermodels.Role]("Accounts")
	return &RoleController{
		GenericController: *generiController,
	}
}

type AccountController struct {
	GenericController[usermodels.Account]
}

func NewAccountController() *AccountController {
	generiController := NewGenericController[usermodels.Account]("Roles")
	return &AccountController{
		GenericController: *generiController,
	}
}

func (controller *AccountController) AssignRole() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Account, usermodels.Role](controller.RelationName)
}

type AuthorController struct {
	GenericController[usermodels.Author]
}

func NewAuthorController() *AuthorController {
	generiController := NewGenericController[usermodels.Author]("Publishers")
	return &AuthorController{
		GenericController: *generiController,
	}
}

func (controller *AuthorController) AssignPublisher() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher](controller.RelationName)
}

type CustomerController struct {
	GenericController[usermodels.Customer]
}


func NewCustomerController() *CustomerController {
	generiController := NewGenericController[usermodels.Customer]("Purchase")
	return &CustomerController{
		GenericController: *generiController,
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


