package migrations

import (
	database "api/bookstoreApi/database"
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
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
		bookmodels.Book{},
		bookmodels.HardCoverFormat{},
		bookmodels.DigitalFormat{},
		paymentmodels.Payment{},
		paymentmodels.Order{},
		paymentmodels.OrderDetails{},
		paymentmodels.Review{},
		paymentmodels.CreditCard{},
		paymentmodels.BankAccount{},
	)
}
