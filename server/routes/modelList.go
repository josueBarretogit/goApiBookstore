package routes

import "api/bookstoreApi/controllers"

type ModelListFormat struct {
	ModelName  string
	Controller controllers.IController
}

func ModelList() []ModelListFormat {
	modelList := []ModelListFormat{
		{ModelName: "role", Controller: controllers.NewRoleController()},
		{ModelName: "author", Controller: controllers.NewAuthorController()},
		{ModelName: "account", Controller: controllers.NewAccountController()},
		{ModelName: "customer", Controller: controllers.NewCustomerController()},
		{ModelName: "publisher", Controller: controllers.NewPublisherController()},
		{ModelName: "book", Controller: controllers.NewBookController()},
		{ModelName: "bookFormat", Controller: controllers.NewBookFormatController()},
		{ModelName: "hardCoverFormat", Controller: controllers.NewHardCoverFormatController()},
		{ModelName: "digitalFormat", Controller: controllers.NewDigitalFormatController()},
		{ModelName: "purchase", Controller: controllers.NewPurchaseController()},
		{ModelName: "purchaseDetails", Controller: controllers.NewPurchaseDetailsController()},
		{ModelName: "review", Controller: controllers.NewReviewController()},
		{ModelName: "paymentMethod", Controller: controllers.NewPaymentMethodController()},
		{ModelName: "creditCard", Controller: controllers.NewCreditCardController()},
		{ModelName: "bankAccount", Controller: controllers.NewBankAccountControler()},
	}
	return modelList
}




