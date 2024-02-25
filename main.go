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


	accountController := controllers.NewAccountController()
	roleController := controllers.NewRoleController()
	publisherController := controllers.NewPublisherController()
	authorController := controllers.NewAuthorController()
	customerController := controllers.NewCustomerController()

	r := gin.Default()

	routes.SetupRoutes("role", roleController, r)
	routes.SetupRoutes("publisher", publisherController, r)
	routes.SetupRoutes("account", accountController, r)
	routes.SetupRoutes("author", authorController, r)
	routes.SetupRoutes("customer", customerController, r)

	r.PUT("account/AssignManyToManyRelation/:id", accountController.AssignRole())
	r.Run()
}
