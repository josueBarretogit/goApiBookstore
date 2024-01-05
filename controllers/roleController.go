package controllers

import (
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repoService database.IRepository
}

func NewRoleController(repo database.IRepository) *RoleController {
	return &RoleController{
		repoService: repo,
	}
}

func (roleController RoleController) CreateRole(c *gin.Context) {
	var role usermodels.Role
	c.BindJSON(&role)
	err, createdRole := roleController.repoService.Create(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating role",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "entity created successfully",
		"entity":  createdRole,
	})

	return
}

func (roleController RoleController) FindAllRole(c *gin.Context) {
	var roles []usermodels.Role

	error := roleController.repoService.Find(&roles)

	if error != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred when getting roles",
		})
		return
	}

	c.JSON(200, gin.H{
		"entities": roles,
	})

}
