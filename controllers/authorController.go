package controllers

import (
	"api/bookstoreApi/consts"
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	GenericController[usermodels.Author]
}

func (controller *AuthorController) AssignPublisher() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher](controller.RelationName, consts.AuthorModelName)
}

func (controller *AuthorController) AssignBook() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Author, usermodels.Publisher]("Books", consts.AuthorModelName)
}
