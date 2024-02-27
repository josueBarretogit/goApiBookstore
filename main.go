package main

import (
	"api/bookstoreApi/config"
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

	r := gin.Default()

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	r.Run()
}
