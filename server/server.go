package server

import (
	"api/bookstoreApi/server/routes"

	"github.com/gin-gonic/gin"
)



func SetupServer() *gin.Engine {
	r := gin.Default()

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	routes.SetupRoutesAccount(r)
	routes.SetupRoutesAuthor(r)
	routes.SetupRoutesPublisher(r)
	return r
}
