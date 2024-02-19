package controllers

import (
	usermodels "api/bookstoreApi/models/userModels"

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
