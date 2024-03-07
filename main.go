package main

import (
	"os"

	"api/bookstoreApi/config"
	"api/bookstoreApi/database"
	"api/bookstoreApi/database/migrations"
	"api/bookstoreApi/server"
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

	r := server.SetupServer()
	errRoute := r.Run()
	if errRoute != nil {
		panic(errRoute.Error())
	}
}
