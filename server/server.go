package server

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/database"
	"api/bookstoreApi/server/routes"
	"flag"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	errEnv := config.LoadEnv()

	if errEnv != nil {
		panic(errEnv.Error())
	}

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}

	if flag.Lookup("test.v") != nil {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.Use(gin.Recovery())

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	routes.SetupRoutesAccount(r)
	routes.SetupRoutesAuthor(r)
	routes.SetupRoutesPublisher(r)
	return r
}
