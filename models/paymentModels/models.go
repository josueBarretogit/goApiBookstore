package paymentmodels

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	DatePurchased   time.Time         `json:"date_purchased,omitempty"`
	ChargedDelivery bool              `json:"charged_delivery,omitempty"`
	UuidPurchase    string            `json:"uuid_purchase,omitempty"`
	AddressShipTo   string            `json:"address_ship_to,omitempty"`
	TotalPrice      float32           `json:"total_price,omitempty"`
	deliveredDate   time.Time         `json:"delivered_date,omitempty"`
	CustomerID      *uint              `json:"customer_id,omitempty"`
	PurchaseDetails []PurchaseDetails `json:"purchase_details,omitempty"`
	PaymentMethodID *uint              `json:"payment_method_id,omitempty"`
}

type PurchaseDetails struct {
	gorm.Model
	Amount     int  `json:"amount,omitempty"`
	PurchaseID *uint `json:"purchase_id,omitempty"`
	BookID     *uint `json:"book_id,omitempty"`
}

type Review struct {
	gorm.Model
	Rating     int    `json:"rating,omitempty"`
	Title      string `json:"title,omitempty"`
	BodyReview string `json:"body_review,omitempty"`
	BookID     *uint   `json:"book_id,omitempty"`
	CustomerID *uint `json:"customer_id,omitempty"`
}

type PaymentMethod struct {
	gorm.Model
	Name        string      `json:"name,omitempty"`
	Purchases   []Purchase  `json:"purchases,omitempty"`
	BankAccount BankAccount `json:"bank_account,omitempty"`
	CreditCard  CreditCard  `json:"credit_card,omitempty"`
}

type BankAccount struct {
	gorm.Model
	BankProvider    string `json:"bank_provider,omitempty"`
	BankNumber      string `json:"bank_number,omitempty"`
	PaymentMethodID *uint   `json:"payment_method_id,omitempty"`
}

type CreditCard struct {
	gorm.Model
	CardNumber      string    `json:"card_number,omitempty"`
	ExpirationDate  time.Time `json:"expiration_date,omitempty"`
	SecurityCodeCvv string    `json:"security_code_cvv,omitempty"`
	PaymentMethodID *uint      `json:"payment_method_id,omitempty"`
}
