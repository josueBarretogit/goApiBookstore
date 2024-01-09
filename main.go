package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	roleController := controllers.NewRoleController(&database.GORMDbRepository{})

	r := gin.Default()

	r.POST("role/save", roleController.CreateRole)
	r.PUT("role/update/:id", roleController.UpdateRole)
	r.DELETE("role/delete/:id", roleController.DeleteRole)
	r.GET("role/findall", roleController.FindAllRole)
	r.GET("role/findby/:id", roleController.FindBy)

	r.Run()
}
