package consts

import (
	"os"
	"path/filepath"
)

const (
	RoleModelName            = "role"
	AuthorModelName          = "author"
	GenreModelName          = "genre"
	AccountModelName         = "account"
	CustomerModelName        = "customer"
	PublisherModelName       = "publisher"
	BookModelName            = "book"
	HardcoverFormatModelName = "hardcoverFormat"
	DigitalFormatModelName   = "digitalFormat"
	AudioBookFormatModelName = "audioBookFormat"
	OrderModelName           = "order"
	OrderDetailsModelName    = "orderDetails"
	ReviewModelName          = "review"
	PaymentModelName         = "payment"
	CreditCardModelName      = "creditCard"
	BankAccountModelName     = "bankAccount"
	RouteFindAll             = "/findall"
	RouteFindById            = "/findby/:id"
	RouteCreate              = "/save"
	RouteUpdate              = "/update/:id"
	RouteDelete              = "/delete/:id"
	RouteBookImageUpload              = "/uploadBookImages/:id"
)

func GetRootDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
