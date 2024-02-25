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

type RoleController struct {
	GenericController[usermodels.Role]
}

type AccountController struct {
	GenericController[usermodels.Account]
}

type AuthorController struct {
	GenericController[usermodels.Author]
}

type CustomerController struct {
	GenericController[usermodels.Customer]
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
