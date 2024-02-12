package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

  roleController := &controllers.RoleController{}
  publisherController := &controllers.PublisherController{}
  accountController := &controllers.PublisherController{}


	r := gin.Default()
  routes.SetupRoutes("role", roleController, r)
  routes.SetupRoutes("publisher", publisherController, r)

	r.Run()
}
