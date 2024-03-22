package usermodels

import (
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}

type Book struct {
	gorm.Model
	Title           string                        `json:"title,omitempty"`
	CoverPhotoUrl   string                        `json:"coverPhotoUrl,omitempty"`
	Description     string                        `json:"description,omitempty"`
	Rating          any                          `json:"rating,omitempty" gorm:"type:int8[]"`
	PublicationDate time.Time                     `json:"publicationDate,omitempty"`
	ISBN            string                        `json:"isbn,omitempty"`
	Ranking         string                        `json:"ranking,omitempty"`
	Authors         []*Author                     `gorm:"many2many:author_book;" json:"authors,omitempty"`
	Languages       []*Language                   `gorm:"many2many:language_book;" json:"languages,omitempty"`
	OrderDetails    []*paymentmodels.OrderDetails `json:"purchase_details,omitempty"`
	GenreID         uint                          `json:"idGenre,omitempty"`
	Genre           Genre                         `json:"genre,omitempty"`
	AudioFormat     bookmodels.AudioBookFormat    `json:"audioFormat,omitempty"`
	HardCoverFormat bookmodels.HardCoverFormat    `json:"hardCover_format,omitempty"`
	DigitalFormat   bookmodels.DigitalFormat      `json:"digitalFormat,omitempty"`
}

type Language struct {
	gorm.Model
	Name  string  `json:"name,omitempty"`
	Books []*Book `gorm:"many2many:language_book;" json:"books,omitempty"`
}

type Author struct {
	gorm.Model
	Name              string       `json:"name" `
	Lastname          string       `json:"lastname" `
	About             *string      `json:"about" `
	ProfilePictureUrl *string      `json:"profilePictureUrl" `
	AccountID         uint         `json:"idAccount" `
	Account           Account      `json:"account"`
	Books             []*Book      `gorm:"many2many:author_book;" `
	Publishers        []*Publisher `gorm:"many2many:author_publisher;" `
}

type Customer struct {
	gorm.Model
	Name              string                  `json:"name,omitempty"`
	Lastname          string                  `json:"lastname" `
	Document          string                  `json:"document,omitempty"`
	PhoneNumber       string                  `json:"phone_number,omitempty"`
	ProfilePictureUrl *string                 `json:"profilePictureUrl,omitempty"`
	AccountID         uint                    `json:"idAccount" `
	Account           Account                 `json:"account" `
	Reviews           []*paymentmodels.Review `json:"reviews"`
}

type Role struct {
	gorm.Model
	Rolename string `json:"rolename" `
}

type Account struct {
	gorm.Model
	Username    string                       `json:"username" `
	Password    string                       `json:"password" `
	Roles       []*Role                      `gorm:"many2many:role_accounts;" json:"roles"`
	BankAccount []*paymentmodels.BankAccount `json:"bank_account" `
	CreditCard  []*paymentmodels.CreditCard  `json:"credit_card" `
}

type Publisher struct {
	gorm.Model
	PublisherName string    `json:"publisherName" `
	Authors       []*Author `gorm:"many2many:author_publisher;" json:"publishers" `
}

type CustomerAddress struct {
	gorm.Model
	Address    string `json:"address" `
	CustomerID uint   `json:"idCustomer"`
}
