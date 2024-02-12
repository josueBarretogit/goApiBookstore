package controllers

import (
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
)

type IController interface {
	FindAll() gin.HandlerFunc
	Create()  gin.HandlerFunc
	Update()  gin.HandlerFunc
	FindOneBy()  gin.HandlerFunc
	Delete()  gin.HandlerFunc
}

type PublisherController struct {
  GenericController[usermodels.Publisher]
}

type RoleController struct {
  GenericController[usermodels.Role]
}





