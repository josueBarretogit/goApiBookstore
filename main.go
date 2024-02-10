package main

import (
	"api/bookstoreApi/config"
	"api/bookstoreApi/controllers"
	"api/bookstoreApi/database"
	"api/bookstoreApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	dbErr := database.ConnectToDB()

	if dbErr != nil {
		panic("Couldnt connect to db")
	}



	r := gin.Default()
	r.POST("prueba/save", controllers.TestCreate[models.Prueba](models.Prueba{}))
	r.GET("prueba/findall", controllers.TestList[models.Prueba](models.Prueba{}))


	r.Run()
}
