package migrations

import (
	database "api/bookstoreApi/database"
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
)

func Migrate() error {
	error := database.ConnectToDB()
	if error != nil {
		panic("Error connecting to db")
	}
	return database.DB.AutoMigrate(
		usermodels.Role{},
		usermodels.Account{},
		usermodels.Author{},
		usermodels.Customer{},
		usermodels.Publisher{},
		usermodels.Genre{},
		usermodels.Language{},
		usermodels.Book{},
		bookmodels.HardCoverFormat{},
		bookmodels.DigitalFormat{},
		bookmodels.AudioBookFormat{},
		paymentmodels.Payment{},
		paymentmodels.Order{},
		paymentmodels.OrderDetails{},
		paymentmodels.Review{},
		paymentmodels.CreditCard{},
		paymentmodels.BankAccount{},
	)
}
