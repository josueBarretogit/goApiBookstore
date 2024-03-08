package main

import (
	"os"

	"api/bookstoreApi/database/migrations"
	"api/bookstoreApi/server"
)

func main() {
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
