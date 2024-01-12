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

	roleController := controllers.NewRoleController(&database.GORMDbRepository{})

	r := gin.Default()
	routes.SetupRoutes("role", roleController, r)
	routes.SetupRoutes("account", roleController, r)

	r.Run()
}
