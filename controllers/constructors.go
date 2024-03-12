package controllers

import (
	"api/bookstoreApi/consts"
	bookmodels "api/bookstoreApi/models/bookModels"
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
)

func NewPublisherController() *PublisherController {
	generiController := NewGenericController[usermodels.Publisher]("Authors", consts.PublisherModelName)
	return &PublisherController{
		GenericController: *generiController,
	}
}

func NewRoleController() *RoleController {
	generiController := NewGenericController[usermodels.Role]("Accounts", consts.RoleModelName)
	return &RoleController{
		GenericController: *generiController,
	}
}

func NewAccountController() *AccountController {
	generiController := NewGenericController[usermodels.Account]("Roles", consts.AccountModelName)
	return &AccountController{
		GenericController: *generiController,
	}
}

func NewAuthorController() *AuthorController {
	generiController := NewGenericController[usermodels.Author]("Publishers", consts.AuthorModelName)
	return &AuthorController{
		GenericController: *generiController,
	}
}

func NewCustomerController() *CustomerController {
	generiController := NewGenericController[usermodels.Customer]("Purchase", consts.CustomerModelName)
	return &CustomerController{
		GenericController: *generiController,
	}
}

func NewBookController() *BookController {
	generiController := NewGenericController[usermodels.Book]("Author", consts.BookModelName)
	return &BookController{
		GenericController: *generiController,
	}
}

func NewHardCoverFormatController() *HardCoverFormatController {
	generiController := NewGenericController[bookmodels.HardCoverFormat]("", consts.HardcoverFormatModelName)
	return &HardCoverFormatController{
		GenericController: *generiController,
	}
}

func NewDigitalFormatController() *DigitalFormatController {
	generiController := NewGenericController[bookmodels.DigitalFormat]("", consts.DigitalFormatModelName)
	return &DigitalFormatController{
		GenericController: *generiController,
	}
}

func NewAudioBookFormatController() *AudioBookFormatController {
	generiController := NewGenericController[bookmodels.AudioBookFormat]("", consts.AudioBookFormatModelName)
	return &AudioBookFormatController{
		GenericController: *generiController,
	}
}

func NewOrderController() *OrderController {
	generiController := NewGenericController[paymentmodels.Order]("", consts.OrderModelName)
	return &OrderController{
		GenericController: *generiController,
	}
}

func NewOrderDetailsController() *PurchaseDetailsController {
	generiController := NewGenericController[paymentmodels.OrderDetails]("", consts.OrderDetailsModelName)
	return &PurchaseDetailsController{
		GenericController: *generiController,
	}
}

func NewReviewController() *ReviewController {
	generiController := NewGenericController[paymentmodels.Review]("", consts.ReviewModelName)
	return &ReviewController{
		GenericController: *generiController,
	}
}

func NewPaymentController() *PaymentController {
	generiController := NewGenericController[paymentmodels.Payment]("", consts.PaymentModelName)
	return &PaymentController{
		GenericController: *generiController,
	}
}

func NewCreditCardController() *CreditCardController {
	generiController := NewGenericController[paymentmodels.CreditCard]("", consts.CreditCardModelName)
	return &CreditCardController{
		GenericController: *generiController,
	}
}

func NewBankAccountControler() *BankAccountController {
	generiController := NewGenericController[paymentmodels.BankAccount]("", consts.BankAccountModelName)
	return &BankAccountController{
		GenericController: *generiController,
	}
}

func NewImageController(module string) *ImageController {
	return &ImageController{
		Module: module,
	}
}

func NewGenreController() *GenreController {
	generiController := NewGenericController[usermodels.Genre]("Books", consts.GenreModelName)
	return &GenreController{
		GenericController: *generiController,
	}
}
