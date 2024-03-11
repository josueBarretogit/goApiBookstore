package controllers

import (

	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"

)


type RoleController struct {
	GenericController[usermodels.Role]
}

type HardCoverFormatController struct {
	GenericController[bookmodels.HardCoverFormat]
}

type DigitalFormatController struct {
	GenericController[bookmodels.DigitalFormat]
}

type AudioBookFormatController struct {
	GenericController[bookmodels.AudioBookFormat]
}
type OrderController struct {
	GenericController[paymentmodels.Order]
}

type PaymentController struct {
	GenericController[paymentmodels.Payment]
}

type PurchaseDetailsController struct {
	GenericController[paymentmodels.OrderDetails]
}

type CreditCardController struct {
	GenericController[paymentmodels.CreditCard]
}

type BankAccountController struct {
	GenericController[paymentmodels.BankAccount]
}

type ReviewController struct {
	GenericController[paymentmodels.Review]
}



