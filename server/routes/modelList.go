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
		{ModelName: "hardCoverFormat", Controller: controllers.NewHardCoverFormatController()},
		{ModelName: "digitalFormat", Controller: controllers.NewDigitalFormatController()},
		{ModelName: "order", Controller: controllers.NewOrderController()},
		{ModelName: "orderDetails", Controller: controllers.NewOrderDetailsController()},
		{ModelName: "review", Controller: controllers.NewReviewController()},
		{ModelName: "payment", Controller: controllers.NewPaymentController()},
		{ModelName: "creditCard", Controller: controllers.NewCreditCardController()},
		{ModelName: "bankAccount", Controller: controllers.NewBankAccountControler()},
	}
	return modelList
}
