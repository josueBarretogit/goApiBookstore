package migrations

import (
	database "api/bookstoreApi/database"
	usermodels "api/bookstoreApi/models/userModels"
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
	)
}
