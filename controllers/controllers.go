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
	AssignManyToManyRelation() gin.HandlerFunc
}

type PublisherController struct {
	GenericController[usermodels.Publisher, usermodels.Account]
}

type RoleController struct {
	GenericController[usermodels.Role, usermodels.Account]
}

type AccountController struct {
	GenericController[usermodels.Account, usermodels.Role]
}

type AuthorController struct {
	GenericController[usermodels.Author, usermodels.Account]
}

type CustomerController struct {
	GenericController[usermodels.Customer, usermodels.Account]
}
