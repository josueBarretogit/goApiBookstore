package controllers

import (
	"net/http"

	"api/bookstoreApi/database"

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
}

func NewGenericController[T interface{}](relation string) *GenericController[T] {
	return &GenericController[T]{
		RelationName: relation,
	}
}

func (controller *GenericController[T]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model T

		errPayload := c.BindJSON(&model)

		if errPayload != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Received bad data",
				"details": errPayload,
			})
			return
		}

		err := database.DB.Create(&model)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
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
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
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
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
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
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		errJson := c.BindJSON(&modelData)
		if errJson != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errJson.Error,
			})
			return
		}

		errDatabase := database.DB.Model(&modelToUpdate).Updates(&modelData)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": errDatabase.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"updated": modelToUpdate,
		})
	}
}

func (controller *GenericController[T]) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelToDelete T
		id := c.Params.ByName("id")
		err := database.DB.First(&modelToDelete, id)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		errDatabase := database.DB.Delete(&modelToDelete)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": errDatabase.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"deleted": modelToDelete,
		})
	}
}

