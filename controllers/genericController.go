package controllers

import (
	"api/bookstoreApi/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type GenericController[T interface{}] struct{}

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
		return
	}
}

func (controller *GenericController[T]) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var models []T
		err := database.DB.Preload(clause.Associations).Find(&models)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"models": models,
		})
		return
	}
}

func (controller *GenericController[T]) FindOneBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model T

		id := c.Params.ByName("id")

		err := database.DB.Limit(1).Find(&model, id)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		if &model == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "That model does not exist",
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
		c.BindJSON(&modelData)
		err = database.DB.Model(&modelToUpdate).Updates(&modelData)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"updated": modelToUpdate,
		})
		return
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
		err = database.DB.Delete(&modelToDelete)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"dbError": err.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"deleted": modelToDelete,
		})
		return
	}
}
