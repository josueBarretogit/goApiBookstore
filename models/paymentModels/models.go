package paymentmodels

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ChargedDelivery bool           `json:"chargedDelivery,omitempty"`
	UuidOrder       string         `json:"uuidPurchase,omitempty"`
	AddressShipTo   string         `json:"addressShip_to,omitempty"`
	DeliveredDate   time.Time      `json:"deliveredDate,omitempty"`
	CustomerID      uint           `json:"idCustomer,omitempty"`
	PaymentID       uint           `json:"idPayment,omitempty"`
	OrderDetailes   []OrderDetails `json:"orderDetails,omitempty"`
}

type OrderDetails struct {
	gorm.Model
	Amount  int  `json:"amount,omitempty"`
	OrderID uint `json:"idOrder,omitempty"`
	BookID  uint `json:"idBook,omitempty"`
}

type Review struct {
	gorm.Model
	Rating     int    `json:"rating,omitempty"`
	Title      string `json:"title,omitempty"`
	BodyReview string `json:"body,omitempty"`
	BookID     uint   `json:"idBook,omitempty"`
	CustomerID uint   `json:"idCustomer,omitempty"`
	PurchaseID *uint  `json:"idPurchase,omitempty"`
}

type Payment struct {
	gorm.Model
	DatePayment   time.Time `json:"datePayment,omitempty"`
	TotalPayed    float64   `json:"totalPrice,omitempty"`
	BankAccountID *uint     `json:"idBankAccount,omitempty"`
	CreditCardID  *uint     `json:"idCreditCard,omitempty"`
}

type BankAccount struct {
	gorm.Model
	BankProvider string `json:"bank_provider,omitempty"`
	BankNumber   string `json:"bank_number,omitempty"`
	AccountID    uint   `json:"account_id,omitempty"`
}

type CreditCard struct {
	gorm.Model
	CardNumber      string    `json:"cardNumber,omitempty"`
	ExpirationDate  time.Time `json:"expirationDate,omitempty"`
	SecurityCodeCvv string    `json:"securityCodeCvv,omitempty"`
	AccountID       uint      `json:"idAccount,omitempty"`
}
