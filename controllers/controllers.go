package controllers

import (
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
)

type RoleController struct {
	GenericController[usermodels.Role]
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
