package controllers

import (
	usermodels "api/bookstoreApi/models/userModels"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var role usermodels.Role
	c.BindJSON(&role)
	err := Models.AddNewBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
