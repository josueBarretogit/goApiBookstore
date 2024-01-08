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

type roleCreateDto struct {
	Rolename string `json:"rolename"`
}

func (roleController *RoleController) CreateRole(c *gin.Context) {
	var role usermodels.Role
	err := c.Bind(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "received bad format",
			"format":   role,
		})
		return
	}
	errCreation := roleController.repoService.Create(&role)
	if errCreation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating role",
			"error":   errCreation,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "entity created successfully",
		"entity":  role,
	})

	return
}

func (roleController *RoleController) FindAllRole(c *gin.Context) {
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

func (roleController *RoleController) UpdateRole(c *gin.Context) {
	var role usermodels.Role
	err := c.Bind(&role)
	id := c.Param("id")
	newRole := usermodels.Role{Rolename: role.Rolename}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "received bad format",
			"format":   role,
		})
		return
	}
	errCreation := roleController.repoService.Create(&newRole)
	if errCreation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating role",
			"error":   errCreation,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "entity created successfully",
		"entity":  newRole,
	})

	return
}
