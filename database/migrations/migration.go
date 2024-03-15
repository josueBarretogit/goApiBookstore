package migrations

import (
	database "api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
)

func Migrate() error {
	error := database.ConnectToDB()
	if error != nil {
		panic("Error connecting to db")
	}
	return database.DB.AutoMigrate(
		&usermodels.Language{},
	)
}
