package server

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/database"
	"api/bookstoreApi/server/routes"
	"flag"
	"time"
 "github.com/gin-contrib/gzip"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET" , "DELETE"},
		AllowHeaders:     []string{"Origin", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		
		MaxAge: 12 * time.Hour,
		}))

  r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.StaticFS("/assets" , gin.Dir("public", false))

	for _, modelFormat := range routes.ModelList() {
		routes.SetupRoutes(modelFormat.ModelName, modelFormat.Controller, r)
	}

	routes.SetupRoutesAccount(r)
	routes.SetupRoutesAuthor(r)
	routes.SetupRoutesPublisher(r)
	routes.SetupRoutesBookRoutes(r)
	routes.SetupRoutesCustomer(r)
	routes.SetupRoutesGenre(r)
	return r
}
