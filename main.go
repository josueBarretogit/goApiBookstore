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

	accountGenericController := controllers.GenericController[usermodels.Account, usermodels.Role]{
		RelationName: "Roles",
	}

	accountController := &controllers.AccountController{
		GenericController: accountGenericController,
	}

	roleController := &controllers.RoleController{}
	publisherController := &controllers.PublisherController{}
	authorController := &controllers.AuthorController{}
	customerController := &controllers.CustomerController{}

	r := gin.Default()

	routes.SetupRoutes("role", roleController, r)
	routes.SetupRoutes("publisher", publisherController, r)
	routes.SetupRoutes("account", accountController, r)
	routes.SetupRoutes("author", authorController, r)
	routes.SetupRoutes("customer", customerController, r)

	r.Run()
}
