package usermodels

import (
	paymentmodels "api/bookstoreApi/models/paymentModels"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name              string            `json:"name" `
	Lastname          string            `json:"lastname" `
	About             string            `json:"about" `
	ProfilePictureUrl string            `json:"profilePictureUrl" `
	AccountID         uint              `json:"accountid" `
	Account           Account           `json:"account"`
	Publishers []*Publisher `gorm:"many2many:author_publisher;" `
}

type Customer struct {
	gorm.Model
	Name              string  `json:"name,omitempty"`
	Document          string  `json:"document,omitempty"`
	Address           string  `json:"address,omitempty"`
	PhoneNumber       string  `json:"phone_number,omitempty"`
	ProfilePictureUrl string  `json:"profile_picture_url,omitempty"`
	AccountID         uint    `json:"accountid" `
	Account           Account `json:"account" `
	Reviews []paymentmodels.Review `json:"reviews"` 
}

type Role struct {
	gorm.Model
	Rolename string `json:"rolename" `
	Accounts []*Account `gorm:"many2many:role_accounts;"`
}

type Account struct {
	gorm.Model
	Username string `json:"username" `
	Password string `json:"password" `
	Roles []*Role `gorm:"many2many:role_accounts;"`
}


type Publisher struct {
	gorm.Model
	PublisherName   string `json:"publisherName" `
	Authors []*Author `gorm:"many2many:author_publisher;" `
}

