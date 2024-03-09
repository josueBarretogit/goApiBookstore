package controllers

import (
	"api/bookstoreApi/consts"
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

func NewAudioBookFormatController() *AudioBookFormatController {
	generiController := NewGenericController[bookmodels.AudioBookFormat]("")
	return &AudioBookFormatController{
		GenericController: *generiController,
	}
}

func NewOrderController() *OrderController {
	generiController := NewGenericController[paymentmodels.Order]("")
	return &OrderController{
		GenericController: *generiController,
	}
}

func NewOrderDetailsController() *PurchaseDetailsController {
	generiController := NewGenericController[paymentmodels.OrderDetails]("")
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

func NewPaymentController() *PaymentController {
	generiController := NewGenericController[paymentmodels.Payment]("")
	return &PaymentController{
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

func NewBookImageController(directoryToStoreImages string) *ImageController {
	return &ImageController{
		DirectoryToStoreImagesPath: directoryToStoreImages,
		Module: consts.BookModelName,
	}
}
