package paymentmodels

import (
	bookmodels "api/bookstoreApi/models/bookModels"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	DatePurchased   time.Time
	ChargedDelivery bool
	UuidPurchase    string
	AddressShipTo   string
	TotalPrice      float32
	deliveredDate   time.Time
	CustomerID      uint
	PurchaseDetails []PurchaseDetails
	PaymentMethodID uint
}

type PurchaseDetails struct {
	gorm.Model
	Amount     int
	Book       bookmodels.Book
	Review     Review
	PurchaseID uint
}

type Review struct {
	gorm.Model
	Rating            int
	Title             string
	BodyReview        string
	PurchaseDetailsID uint
}

type PaymentMethod struct {
	gorm.Model
	Name        string
	Purchases   []Purchase
	BankAccount BankAccount
	CreditCard  CreditCard
}

type BankAccount struct {
	gorm.Model
	BankProvider    string
	BankNumber      string
	PaymentMethodID uint
}

type CreditCard struct {
	gorm.Model
	CardNumber      string
	ExpirationDate  time.Time
	SecurityCodeCvv string
	PaymentMethodID uint
}
