package main

import (
	"os"

	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/database/migrations"
	"api/bookstoreApi/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	errEnv := config.LoadEnv()

	if errEnv != nil {
		panic(errEnv.Error())
	}

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	if os.Getenv("MIGRATE") != "" {
		errMigration := migrations.Migrate()
		if errMigration != nil {
			panic(errMigration.Error())
		}
		return
	}

	r := gin.Default()

	routes.SetupRoutesAccount(r)

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	r.PUT("author/assignPublisher/:id", controllers.NewAuthorController().AssignPublisher())
	r.PUT("publisher/assignAuthor/:id", controllers.NewPublisherController().AssignAuthor())

	errRoute := r.Run()
	if errRoute != nil {
		panic(errRoute.Error())
	}
}
