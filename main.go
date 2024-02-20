package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/database/migrations"
	"api/bookstoreApi/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	migrations.Migrate()

	roleController := &controllers.RoleController{}
	accountController := &controllers.AccountController{}
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
