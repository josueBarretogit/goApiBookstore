package migrations

import (
	database "api/bookstoreApi/database"
	"api/bookstoreApi/models"
	usermodels "api/bookstoreApi/models/userModels"
	"log"
)

func Migrate() {
	error := database.ConnectToDB()
	if error != nil {
		panic("Error connecting to db")
	}
	database.DB.AutoMigrate(
		usermodels.Role{},
		usermodels.Account{},
		usermodels.Author{},
		usermodels.Customer{},
		usermodels.Publisher{},
		usermodels.PublisherAuthor{},
	)
}

func MigrateTest() {
	error := database.ConnectToDB()
	if error != nil {
		log.Fatal("Something happened when migrating")
	}
	database.DB.AutoMigrate(
		&models.Prueba{},
	)
}
