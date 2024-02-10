package usermodels

import (
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name              string
	Lastname          string
	About             string
	ProfilePictureUrl string
	AccountID         uint
	Account           Account
	PublisherAuthor   []PublisherAuthor
	Book              []bookmodels.Book
}

type Customer struct {
	gorm.Model
	Name              string
	Lastname          string
	Document          string
	Address           string
	PhoneNumber       string
	ProfilePictureUrl string
	AccountID         uint
	Account           Account
	Purchases         []paymentmodels.Purchase
}

type Account struct {
	gorm.Model
	Username string
	Password string
	RoleID   uint
	Role     Role
}

type Role struct {
	gorm.Model
	ID       int    `gorm:"AUTO_INCREMENT"`
	Rolename string `json:"rolename" binding:"required"`
}

type Publisher struct {
	gorm.Model
	PublisherName   string `json:"publisherName" binding:"required"`
	PublisherAuthor []PublisherAuthor
}

type PublisherAuthor struct {
	gorm.Model
	PublisherName string
	PublisherID   uint
	AuthorID      uint
}
