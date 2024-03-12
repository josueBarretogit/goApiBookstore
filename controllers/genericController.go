package controllers

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type IController interface {
	FindAll() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	FindOneBy() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type GenericController[T interface{}] struct {
	RelationName string
	ModelName    string
}

func NewGenericController[T interface{}](relation string, modelName string) *GenericController[T] {
	return &GenericController[T]{
		RelationName: relation,
		ModelName:    modelName,
	}
}

func (controller *GenericController[T]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model T

		errPayload := c.BindJSON(&model)

		if errPayload != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeBadData,
				"target":  controller.ModelName,
				"details": errPayload.Error(),
			})
			return
		}

		err := database.DB.Create(&model)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"created": model,
		})
	}
}

func (controller *GenericController[T]) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var models []T
		err := database.DB.Preload(clause.Associations).Order("ID desc").Find(&models)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"models": models,
		})
	}
}

func (controller *GenericController[T]) FindOneBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model T

		id := c.Params.ByName("id")

		err := database.DB.Limit(1).Preload(clause.Associations).Find(&model, id)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"model": model,
		})
	}
}

func (controller *GenericController[T]) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelToUpdate T
		var modelData T

		id := c.Params.ByName("id")
		err := database.DB.First(&modelToUpdate, id)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": err.Error.Error(),
			})
			return
		}
		errJson := c.BindJSON(&modelData)
		if errJson != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": errJson.Error(),
			})
			return
		}

		errDatabase := database.DB.Model(&modelToUpdate).Updates(&modelData)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": errDatabase.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"updated": modelToUpdate,
			"success": true,
		})
	}
}

func (controller *GenericController[T]) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelToDelete T
		id := c.Params.ByName("id")
		err := database.DB.First(&modelToDelete, id)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": err.Error.Error(),
			})
			return
		}
		errDatabase := database.DB.Delete(&modelToDelete)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  controller.ModelName,
				"details": errDatabase.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"deleted": modelToDelete,
			"success": true,
		})
	}
}

func AssignManyToManyRelation[T interface{}, K interface{}](relation string, target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelToUpdate T
		var modelData K

		id := c.Params.ByName("id")
		err := database.DB.First(&modelToUpdate, id)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  target,
				"details": err.Error.Error(),
			})
			return
		}

		errJson := c.BindJSON(&modelData)
		if errJson != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    consts.ErrorCodeBadData,
				"target":  target,
				"details": errJson.Error(),
			})
			return
		}

		errDatabase := database.DB.Model(&modelToUpdate).Association(relation).Append(&modelData)
		if err.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    consts.ErrorCodeDatabase,
				"target":  target,
				"details": errDatabase.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"updated": modelData,
			"success": true,
		})
	}
}
