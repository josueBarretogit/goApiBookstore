package controllers

import (
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
)


type AuthorController struct {
	GenericController[usermodels.Author]
}

func (controller *AuthorController) AssignPublisher() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher](controller.RelationName)
}

func (controller *AuthorController) AssignBook() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher]("Books")
}
