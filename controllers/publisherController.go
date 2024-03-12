package controllers

import (
	"api/bookstoreApi/consts"
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
)

type PublisherController struct {
	GenericController[usermodels.Publisher]
}

func (controller *PublisherController) AssignAuthor() gin.HandlerFunc {
	return AssignManyToManyRelation[usermodels.Publisher, usermodels.Author](controller.RelationName, consts.PublisherModelName)
}
