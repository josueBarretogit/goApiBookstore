package paymentmodels

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ChargedDelivery bool           `json:"charged_delivery,omitempty"`
	UuidOrder       string         `json:"uuid_purchase,omitempty"`
	AddressShipTo   string         `json:"address_ship_to,omitempty"`
	DeliveredDate   time.Time      `json:"delivered_date,omitempty"`
	CustomerID      uint           `json:"customer_id,omitempty"`
	PaymentID       uint           `json:"payment_id,omitempty"`
	OrderDetailes   []OrderDetails `json:"order_details,omitempty"`
}

type OrderDetails struct {
	gorm.Model
	Amount  int  `json:"amount,omitempty"`
	OrderID uint `json:"order_id,omitempty"`
	BookID  uint `json:"book_id,omitempty"`
}

type Review struct {
	gorm.Model
	Rating     int    `json:"rating,omitempty"`
	Title      string `json:"title,omitempty"`
	BodyReview string `json:"body_review,omitempty"`
	BookID     uint   `json:"book_id,omitempty"`
	CustomerID uint   `json:"customer_id,omitempty"`
	PurchaseID *uint  `json:"purchase_id,omitempty"`
}

type Payment struct {
	gorm.Model
	DatePayment   time.Time `json:"date_payment,omitempty"`
	TotalPayed    float64   `json:"total_price,omitempty"`
	BankAccountID *uint     `json:"bank_account,omitempty"`
	CreditCardID  *uint     `json:"credit_card,omitempty"`
}

type BankAccount struct {
	gorm.Model
	BankProvider string `json:"bank_provider,omitempty"`
	BankNumber   string `json:"bank_number,omitempty"`
	AccountID    uint   `json:"account_id,omitempty"`
}

type CreditCard struct {
	gorm.Model
	CardNumber      string    `json:"card_number,omitempty"`
	ExpirationDate  time.Time `json:"expiration_date,omitempty"`
	SecurityCodeCvv string    `json:"security_code_cvv,omitempty"`
	AccountID       uint      `json:"account_id,omitempty"`
}
