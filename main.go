package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/database/migrations"
	"api/bookstoreApi/server/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	if  os.Getenv("MIGRATE") != "" {
		migrations.Migrate()
		return
	}

	r := gin.Default()

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}
	

	r.PUT("account/assignRole/:id" , controllers.NewAccountController().AssignRole())
	r.PUT("author/assignPublisher/:id" , controllers.NewAuthorController().AssignPublisher())
	r.PUT("publisher/assignAuthor/:id" , controllers.NewPublisherController().AssignAuthor())

	r.Run()
}
