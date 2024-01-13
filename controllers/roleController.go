package controllers

import (
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	repositoryService database.IRepository
	model             interface{}
}

func NewRoleController(databaseRepository database.IRepository, model interface{}) *RoleController {
	return &RoleController{
		repositoryService: databaseRepository,
		model:             model,
	}
}

func (roleController *RoleController) Create(c *gin.Context) {
	var role usermodels.Role
	err := c.Bind(&role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "received bad format",
			"format":   role,
		})
		return
	}

	errCreation := roleController.repositoryService.Create(&role)
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

func (roleController *RoleController) FindAll(c *gin.Context) {

	error := roleController.repositoryService.Find(&roleController.model)

	if error != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred when getting roles",
		})
		return
	}

	c.JSON(200, gin.H{
		"entities": roleController.model,
	})

}

func (roleController *RoleController) FindOneBy(c *gin.Context) {
	var role usermodels.Role

	id := c.Param("id")

	findError := roleController.repositoryService.FindOneBy(&role, id)

	if findError != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred when getting roles",
		})
		return
	}

	c.JSON(200, gin.H{
		"entity": role,
	})

}

func (roleController *RoleController) Update(c *gin.Context) {
	var roleData usermodels.Role
	var roleToUpdate usermodels.Role

	err := c.Bind(&roleData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response":   "Received bad json format",
			"conditions": roleData,
		})
		return
	}

	id := c.Param("id")
	findErr := roleController.repositoryService.FindOneBy(&roleToUpdate, id)

	if findErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response":   "Couldnt find model to update",
			"conditions": id,
		})
		return
	}

	errUpdate := roleController.repositoryService.Update(&roleToUpdate, roleData)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error updating role",
			"error":   errUpdate,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "entity updated successfully",
		"entity":  roleToUpdate,
	})

	return
}

func (roleController *RoleController) Delete(c *gin.Context) {

	id := c.Param("id")
	var roleToDelete usermodels.Role

	findErr := roleController.repositoryService.FindOneBy(&roleToDelete, id)

	if findErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"response": "entity not found",
		})
	}

	deleteErr := roleController.repositoryService.Delete(&roleToDelete, id)

	if deleteErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response":   "Couldnt find model to update",
			"conditions": id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "entity deleted successfully",
		"entityDeleted": roleToDelete,
	})

	return
}
