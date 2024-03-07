package main

import (
	"os"

	"api/bookstoreApi/config"
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

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	routes.SetupRoutesAccount(r)
	routes.SetupRoutesAuthor(r)
	routes.SetupRoutesPublisher(r)

	errRoute := r.Run()
	if errRoute != nil {
		panic(errRoute.Error())
	}
}
