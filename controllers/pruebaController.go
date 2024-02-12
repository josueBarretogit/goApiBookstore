package controllers

import (
	"api/bookstoreApi/database"
	"api/bookstoreApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


type PruebaController [T any]  struct {
	repositoryService database.IRepository
	NewModel             func() *T
}


func TestCreate [T any]() gin.HandlerFunc {
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

func TestList[T any]() gin.HandlerFunc {
  return func(c *gin.Context) {
    var pruebas []T
    err := database.DB.Find(&pruebas)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    } 
    c.JSON(http.StatusOK, gin.H{
      "pruebas" : pruebas,
    }) 
    return
  }
}

func TestPut[T any](modelType T) gin.HandlerFunc  {
  return func(c *gin.Context) {
    var pruebaToUpdate models.Prueba
    var prueba models.Prueba

    id := c.Params.ByName("id")
    err := database.DB.First(&prueba, id)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    }
    c.BindJSON(&pruebaToUpdate)
    err = database.DB.Model(&prueba).Updates(*&pruebaToUpdate)
    if err.Error != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "dbError" : err.Error,
      }) 
      return
    } 
    c.JSON(http.StatusOK, gin.H{
      "updated" : pruebaToUpdate,
    }) 
    return
  }
}

func TestDelete(c *gin.Context) {

  var pruebaToDelete models.Prueba

  id := c.Params.ByName("id")
  err := database.DB.First(&pruebaToDelete, id)
  if err.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "dbError" : err.Error,
    }) 
    return
  }
  err = database.DB.Delete(&pruebaToDelete)
  if err.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "dbError" : err.Error,
    }) 
    return
  } 
  c.JSON(http.StatusOK, gin.H{
    "deleted" : pruebaToDelete,
  }) 
  return

}
