package controllers

import (
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
)

func NewPublisherController() *PublisherController {
	generiController := NewGenericController[usermodels.Publisher]("Authors")
	return &PublisherController{
		GenericController: *generiController,
	}
}

func NewRoleController() *RoleController {
	generiController := NewGenericController[usermodels.Role]("Accounts")
	return &RoleController{
		GenericController: *generiController,
	}
}

func NewAccountController() *AccountController {
	generiController := NewGenericController[usermodels.Account]("Roles")
	return &AccountController{
		GenericController: *generiController,
	}
}

func NewAuthorController() *AuthorController {
	generiController := NewGenericController[usermodels.Author]("Publishers")
	return &AuthorController{
		GenericController: *generiController,
	}
}

func NewCustomerController() *CustomerController {
	generiController := NewGenericController[usermodels.Customer]("Purchase")
	return &CustomerController{
		GenericController: *generiController,
	}
}

func NewBookController() *BookController {
	generiController := NewGenericController[bookmodels.Book]("PurchaseDetails")
	return &BookController{
		GenericController: *generiController,
	}
}

func NewBookFormatController() *BookFormatController {
	generiController := NewGenericController[bookmodels.BookFormat]("")
	return &BookFormatController{
		GenericController: *generiController,
	}
}

func NewHardCoverFormatController() *HardCoverFormatController {
	generiController := NewGenericController[bookmodels.HardCoverFormat]("")
	return &HardCoverFormatController{
		GenericController: *generiController,
	}
}

func NewDigitalFormatController() *DigitalFormatController {
	generiController := NewGenericController[bookmodels.DigitalFormat]("")
	return &DigitalFormatController{
		GenericController: *generiController,
	}
}

func NewPurchaseController() *PurchaseController {
	generiController := NewGenericController[paymentmodels.Purchase]("")
	return &PurchaseController{
		GenericController: *generiController,
	}
}

func NewPurchaseDetailsController() *PurchaseDetailsController {
	generiController := NewGenericController[paymentmodels.PurchaseDetails]("")
	return &PurchaseDetailsController{
		GenericController: *generiController,
	}
}

func NewReviewController() *ReviewController {
	generiController := NewGenericController[paymentmodels.Review]("")
	return &ReviewController{
		GenericController: *generiController,
	}
}

func NewPaymentMethodController() *PaymentMethodController {
	generiController := NewGenericController[paymentmodels.PaymentMethod]("")
	return &PaymentMethodController{
		GenericController: *generiController,
	}
}

func NewCreditCardController() *CreditCardController {
	generiController := NewGenericController[paymentmodels.CreditCard]("")
	return &CreditCardController{
		GenericController: *generiController,
	}
}

func NewBankAccountControler() *BankAccountController {
	generiController := NewGenericController[paymentmodels.BankAccount]("")
	return &BankAccountController{
		GenericController: *generiController,
	}
}
