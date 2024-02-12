package controllers

import (
	"api/bookstoreApi/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenericController[T any] struct {
}


func (controller *GenericController[T]) Create() gin.HandlerFunc {
  return func(c *gin.Context) {
    var model T
    c.BindJSON(&model)
    err := database.DB.Create(&model)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    }
    c.JSON(http.StatusOK, gin.H{
      "created" : model,
    })

  }
}

func (controller *GenericController[T]) FindAll() gin.HandlerFunc {
  return func(c *gin.Context) {
    var models []T
    err := database.DB.Find(&models)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    } 
    c.JSON(http.StatusOK, gin.H{
      "models" : models,
    }) 
    return
  }
}


func (controller *GenericController[T]) FindOneBy() gin.HandlerFunc { 
  return func(c *gin.Context) {
  }
}



func (controller *GenericController[T]) Update() gin.HandlerFunc { 
  return func(c *gin.Context) {

    var modelToUpdate T
    var modelData T

    id := c.Params.ByName("id")
    err := database.DB.First(&modelData, id)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    }
    c.BindJSON(&modelToUpdate)
    err = database.DB.Model(&modelData).Updates(*&modelToUpdate)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    } 
    c.JSON(http.StatusOK, gin.H{
      "updated" : modelToUpdate,
    }) 
    return
  }
}



func (controller *GenericController[T]) Delete() gin.HandlerFunc { 
  return func(c *gin.Context) {
  }
}
