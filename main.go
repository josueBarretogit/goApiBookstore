package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
	"api/bookstoreApi/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	var roles []usermodels.Role
	var publishers []usermodels.Publisher
	roleController := controllers.NewRoleController(&database.GORMDbRepository{}, roles)
	publisherController := controllers.NewRoleController(&database.GORMDbRepository{}, publishers)

	r := gin.Default()
	routes.SetupRoutes("role", roleController, r)
	routes.SetupRoutes("publisher", publisherController, r)

	r.Run()
}
