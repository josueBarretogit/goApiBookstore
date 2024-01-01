package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string
	CoverUrl        string
	Description     string
	numPages        int
	Rating          int
	PublicationDate time.Time
	Genre           string
	Language        string
	ISBN            string
	Ranking         string
	Stock           int64
	BookFormat      []BookFormat
}

type BookFormat struct {
	gorm.Model
	FormatName string
	Price      float64
	BookID     uint
}

type DigitalFormat struct {
	gorm.Model
	FormatName string
	Price      float64
	BookID     uint
}

type Author struct {
	gorm.Model
	Name              string
	Lastname          string
	About             string
	ProfilePictureUrl string
	AccountID         uint
	Account           Account
	PublisherAuthor   []PublisherAuthor
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
	Rolename string
}

type Publisher struct {
	gorm.Model
	PublisherName   string
	PublisherAuthor []PublisherAuthor
}

type PublisherAuthor struct {
	gorm.Model
	PublisherName string
	PublisherID   uint
	AuthorID      uint
}
