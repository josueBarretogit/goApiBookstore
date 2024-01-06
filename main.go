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

	r.POST("role/create", roleController.CreateRole)
	r.GET("role/getall", roleController.FindAllRole)
	r.Run()
}
