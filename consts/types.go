package consts

import (
	"os"
	"path/filepath"
)

const (
	RoleModelName            = "roles"
	AuthorModelName          = "authors"
	GenreModelName           = "genres"
	AccountModelName         = "accounts"
	CustomerModelName        = "customers"
	PublisherModelName       = "publishers"
	BookModelName            = "books"
	HardcoverFormatModelName = "hardcoverFormats"
	DigitalFormatModelName   = "digitalFormats"
	AudioBookFormatModelName = "audioBookFormats"
	OrderModelName           = "orders"
	OrderDetailsModelName    = "orderDetails"
	ReviewModelName          = "reviews"
	PaymentModelName         = "payments"
	CreditCardModelName      = "creditCards"
	BankAccountModelName     = "bankAccounts"
	LanguageModelName        = "languages"
	RouteFindAll             = ""
	RouteFindById            = "/:id"
	RouteCreate              = ""
	RouteUpdate              = "/:id"
	RouteDelete              = "/:id"
	RouteBookImageUpload     = "/uploadBookImages/:id"
	RouteBestSellers         = "/bestSellers"
)

func GetRootDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

const (
	ErrorCodeBadData  = "badData"
	ErrorCodeDatabase = "dbError"
	ErrorNotNumber    = "input was not a number"
	ErrorBadDate      = "Received invalid date format"
)

const AVGrating = `(SELECT AVG(ratings) FROM UNNEST(books.rating) ratings) as rating`
