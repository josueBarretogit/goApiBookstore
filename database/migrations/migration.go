package migrations

import (
	database "api/bookstoreApi/database"
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
	"log"
)

func Migrate() {
	error := database.ConnectToDB()
	if error != nil {
		log.Fatal("Something happened when migrating")
	}
	database.DB.AutoMigrate(
		&usermodels.Role{},
		&usermodels.Account{},
		&usermodels.Customer{},
		&usermodels.Author{},
		&bookmodels.Book{},
		&bookmodels.BookFormat{},
		&bookmodels.HardCoverFormat{},
		&bookmodels.DigitalFormat{},
		&usermodels.Publisher{},
		&usermodels.PublisherAuthor{},
		&paymentmodels.PaymentMethod{},
		&paymentmodels.Purchase{},
		&paymentmodels.Review{},
		&paymentmodels.PurchaseDetails{},
		&paymentmodels.CreditCard{},
		&paymentmodels.BankAccount{},
	)

}
