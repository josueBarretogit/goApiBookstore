package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/models"
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

  pruebaController := &controllers.GenericController[models.Prueba]{}
  roleController := &controllers.GenericController[usermodels.Role]{}
  publisherController := &controllers.PublisherController{}


	r := gin.Default()
  routes.SetupRoutes("role", roleController, r)
  routes.SetupRoutes("publisher", publisherController, r)
  routes.SetupRoutes("prueba", pruebaController, r)



	r.Run()
}
